This stupid little thing is to keep me from getting the yaml/base64 decoding k8s secrets all the time
Instead of adding k8s things in here i opted for just adding the function

```
decode () {
  kubectl get secret $1 -o json | k8s-secret-decode
}
```

in my .zshrc

# Example output

```
stan.json
{
    my config content
}

application.toml
[example]
  foo = bar
```
