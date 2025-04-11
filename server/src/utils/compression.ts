import { brotliDecompressSync } from 'zlib';

export const decodeBrotli = (content: string): string => {
    const buffer = Buffer.from(content);
    console.log(buffer);
    const decompressedBuffer = brotliDecompressSync(buffer, { finishFlush: 1 });
    return decompressedBuffer.toString('utf-8');
}


