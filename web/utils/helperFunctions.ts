// function that returns a string from a ReadableStream
async function RSToString(stream: unknown): Promise<string> {
  const decoder = new TextDecoder();
  let result = "";
  const transform = (chunk: Uint8Array) => {
    result += decoder.decode(chunk);
  };
  for await (const chunk of stream as AsyncIterableIterator<Uint8Array>) {
    transform(chunk);
  }

  return result;
}

export { RSToString };
