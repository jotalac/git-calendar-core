# Git Calendar Core
[![Go version](https://img.shields.io/github/go-mod/go-version/git-calendar/core)](./go.mod)
[![Go](https://github.com/git-calendar/core/actions/workflows/go.yaml/badge.svg)](https://github.com/git-calendar/core/actions/workflows/go.yaml)
[![License](https://img.shields.io/github/license/git-calendar/core)](./LICENSE.txt)

A shared core for Git-backed calendar.

Related projects:
- Web client: [git-calendar/web-client](https://github.com/git-calendar/web-client)
- CORS Proxy: [git-calendar/cors-proxy](https://github.com/git-calendar/cors-proxy)

## Build
Build tools are pinned in `go.mod` and invoked through `go tool`.
For Android builds, install the Android SDK and NDK, set `ANDROID_HOME`, and ensure Android API 35 is available.
For iOS builds, use macOS with Xcode installed.

Build a target:
```sh
make build_web
make build_android
make build_ios
```

The Android and iOS builds bind `pkg/mobile`: create the core with `NewCore(dataDir)` (on Android, the app's files directory).
Its errors are JSON objects with `code`, `calendar` and `message` fields.
The web build uses `cmd/wasm`, which wraps `pkg/api` directly.
