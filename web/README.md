# sa web

## Usage

* run `npm run dev` to run the development server in watch mode (equivalent to running `next dev`).
* Run `npm start` to run the production server in watch mode (equivalent to running `next start`).

### Build

`bazel build //next.js:next`, the Bazel equivalent of running `next build`.

The output `.next` folder can be found under `bazel-bin/next.js/.next`.

### Export

`bazel build //next.js:next_export`, the Bazel equivalent of running `next export`.
The output `out` folder can be found under `bazel-bin/next.js/out`.

# Deployment

## Running the container locally

1. `bazel run //:web_tarball` load it
2. `docker run --rm -p 8080:8080 sa:latest`

## Push Image to Registry

You'll need crane first: `go install github.com/google/go-containerregistry/cmd/crane@latest`

Then:

`bazel run //web:push`
