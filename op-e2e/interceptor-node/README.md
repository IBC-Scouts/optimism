### Hola amigo!

Some small things for running e2e's (and some minor pitfalls):

First, add the binary, build and copy in this directory:

```bash
make build-interceptor
```

We also need the peptide binary in the same dir (though I think I can just invoke
its `RootCmd` eventually). In the monomer-poc repo, run:

```bash
make build-peptide
```

Then copy the binary to this dir.

In short, this directory should contain _both_ `interceptor` and `peptide` binaries.

After copying to `interceptor-node` folder, can just invoke a specific e2e with:

```bash
go test -v -run TestDepositTxCreateContract ./...
```

---

Some minor pitfallinos:

Clean up peptide dir that holds genesis data/config etc in `.peptide`. For some
reason passing `--override` to the binary doesn't respect it. Simply:

```bash
rm -rf ~/.peptide
```

In addition to ^, I have the binary just lurking around in the background like a
madman if an e2e fails, just murder it brutally if `ps -aux | grep "interceptor"` shows it.

---

In this dir:

- `interceptor.go` calls our interceptor binary (which will invoke peptide binary).
- `client.go` holds a tiny rpc client we can use to interact with interceptor-node rpc server.
