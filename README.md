# FailToRun

Program periodically checks specified URL (default: `www.google.com`) and if the request fails start specified binary.
Other options like minimum interval, timeout, URL etc. can be specified including negative behavior when it runs the command/script on success response.

## Usages

Mininal usage:\
`failtorun -c "cmd"`

Run on 404:\
`failtorun -u www.mypage9999.org/test -c ./my_batch.sh -r 404 -n 1`

## Options
- `-b` (int)
  - Run command at background. (default 1)
- `-c` (string)
  - Command to run.
- `-m` (int)
  - Maximum number of repeats (0 = unlimited). (default 0)
- `-n` (int)
  - Negator (run at success condition). (default 0)
- `-p` (int)
  - Repeat every N seconds. (default 1)
- `-r` (int)
  - Maximum response code. (default 200)
- `-t` (int)
  - Request timeout in seconds. (default 2)
- `-u` (string)
  - URL to check. (default: "www.google.com")
