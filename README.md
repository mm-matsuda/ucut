# ucut

useful cut

cut's behavior is below.

```shell
% who
user1 console  Nov 26 09:08
user1 ttys000  Nov 26 09:08
user1 ttys001  Nov 26 09:08
user1 ttys002  Nov 26 09:08
% who | cut -d " " -f2,1
user1 console
user1 ttys000
user1 ttys001
user1 ttys002
```

but, I want is below.

```shell
% who | cut -d " " -f2,1
console user1
ttys000 user1
ttys001 user1
ttys002 user1
```
