package sshql

import (
	"fmt"
	"net"
	"os"
	"path/filepath"

	"github.com/goark/errs"
	"golang.org/x/crypto/ssh"
	"golang.org/x/crypto/ssh/agent"
	"golang.org/x/crypto/ssh/knownhosts"
)

// Dialer is authentication provider information.
type Dialer struct {
	Hostname      string `json:"hostname"`
	Port          int    `json:"port"`
	Username      string `json:"username"`
	Password      string `json:"password"`
	PrivateKey    string `json:"privateKey"`
	IgnoreHostKey bool   `json:"IgnoreHostKey"`
	client        *ssh.Client
}

// Connect starts a client connection to the given SSH server.
func (d *Dialer) Connect() error {
	sshConfig := &ssh.ClientConfig{
		User: d.Username,
		Auth: []ssh.AuthMethod{},
		HostKeyCallback: ssh.HostKeyCallback(func(host string, remote net.Addr, pubKey ssh.PublicKey) error {
			if d.IgnoreHostKey {
				return nil
			}
			kh, err := getKnownHosts()
			if err != nil {
				return errs.Wrap(err)
			}
			if err := kh(host, remote, pubKey); err != nil {
				return errs.Wrap(err)
			}
			return nil
		}),
	}

	if conn, err := net.Dial("unix", os.Getenv("SSH_AUTH_SOCK")); err == nil {
		defer conn.Close()
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(agent.NewClient(conn).Signers))
	}

	if d.PrivateKey != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PublicKeysCallback(func() ([]ssh.Signer, error) {
			return getSigners(d.PrivateKey, d.Password)
		}))
	} else if d.Password != "" {
		sshConfig.Auth = append(sshConfig.Auth, ssh.PasswordCallback(func() (string, error) {
			return d.Password, nil
		}))
	}

	sshcon, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", d.Hostname, d.Port), sshConfig)
	if err != nil {
		return errs.Wrap(err)
	}
	d.client = sshcon
	return nil
}

// Dial makes socket connection via SSH tunnel.
func (d *Dialer) Dial(network, address string) (net.Conn, error) {
	return d.client.Dial(network, address)
}

func getSigners(keyfile string, password string) ([]ssh.Signer, error) {
	buf, err := os.ReadFile(keyfile)
	if err != nil {
		return nil, errs.Wrap(err)
	}

	if password != "" {
		k, err := ssh.ParsePrivateKeyWithPassphrase(buf, []byte(password))
		if err != nil {
			return nil, errs.Wrap(err)
		}
		return []ssh.Signer{k}, nil
	}

	k, err := ssh.ParsePrivateKey(buf)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return []ssh.Signer{k}, nil
}

func getKnownHosts() (ssh.HostKeyCallback, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, errs.Wrap(err)
	}
	path := filepath.Join(homeDir, ".ssh", "known_hosts")
	// create known_hosts file if not exist
	if err := func() error {
		f, err := os.OpenFile(path, os.O_CREATE, 0600)
		if err != nil {
			return err
		}
		return f.Close()
	}(); err != nil {
		return nil, errs.Wrap(err)
	}
	kh, err := knownhosts.New(path)
	if err != nil {
		return nil, errs.Wrap(err)
	}
	return kh, nil
}

/**
 * This codes are forked from github.com/mattn/pqssh package.
 */
