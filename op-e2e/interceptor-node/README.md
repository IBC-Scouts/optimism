### Hola amigo!

Some small things for running e2e's (and some minor pitfalls):

First, add the binary, build and copy in this directory:

```bash
make build-interceptor
```

After copying to `interceptor-node` folder, can just invoke a specific e2e with:

```bash
go test -v -run TestDepositTxCreateContract ./...
```

seems like our binary size is 101.37mb which exceeds Github's size limit of 100mb
(according to pre-receive hook) and as such, can't push it on repo atm.

---

Some minor pitfallinos:

Clean up peptide dir that holds genesis data/config etc in `.peptide`. For some
reason passing `--override` to the binary doesn't respect it. Simply:

```bash
rm -rf .peptide
```

In addition to ^, I have the binary just lurking around in the background like a
madman if an e2e fails, just murder it brutally if `ps -aux | grep "interceptor"` shows it.

---

In this dir:

- `interceptor.go` wraps the binary invocations, calls `init` then `seal` then `start`, is used in `setup.go`.
- `client.go` holds a tiny rpc client we can use to interact with interceptor-node rpc server.
