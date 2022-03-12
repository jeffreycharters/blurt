import ioClient from 'socket.io-client';

const ENDPOINT = 'https://letsblurt.duckdns.org/';

const socket = ioClient(ENDPOINT);

export const io = socket;
