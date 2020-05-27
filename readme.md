# go-2b

Interface library and remote operation utilities for the [E-Stim Systems 2B](https://store.e-stim.co.uk/index.php?main_page=index&cPath=23_24) powerbox.
Essentially a cleaner, more reliable rewrite of [2b-utils](https://github.com/DanNixon/2b-utils).

**Note**: The API provided by `2b-server` is unauthenticated, as such do not expose it over the internet.
If you do wish to provide remote access do so via a reverse proxy (such as [Nginx](https://nginx.org/)) with sufficient authentication or via a VPN.

## Disclaimer

This is relatively well tested, however perfect function is not guaranteed.

:warning: Use with caution.

## To Do

- Investigate why channel linking does not work as intended
