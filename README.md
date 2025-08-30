# Blurtonamo Bay

## Make work

1. Run the backend. You'll need go 1.21+ for this.
1. In the `eyes` dir, create a `.env` file with the following entries but set the values to your circumstances:

```env
ORIGIN=http://localhost:5173
PUBLIC_API_ADDRESS=http://localhost:3320
PUBLIC_WS_ADDRESS=ws://localhost:3320
```

1. Run `pnpm install` or `npm install` to in `eyes` to install javascript dependencies.
1. On the same network, run the frontend using `pnpm dev` or whatever.

## TODOs

- [x] Add entgo.io to backend, set up sqlite db
- [x] API endpoint: create user or verify existence
- [x] Think about websockets implementation
