console.log(`Config loaded:\n\n${JSON.stringify(import.meta.env, null, 2)}`);

export const STREAM_BE_BASE_URL = import.meta.env.VITE_STREAM_BE_BASE_URL;
export const STREAM_URL = `${STREAM_BE_BASE_URL}/stream`;
