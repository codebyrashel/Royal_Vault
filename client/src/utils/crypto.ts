// Parameters for PBKDF2
const PBKDF2_ITERATIONS = 150000;
const PBKDF2_HASH = 'SHA-256';
const KEY_LENGTH_BITS = 256; // 256-bit key for AES-GCM

// AES-GCM parameters
const AES_ALGO = 'AES-GCM';
const IV_LENGTH_BYTES = 12; // recommended 12 bytes for GCM

export interface DerivedKeyResult {
  key: CryptoKey;
  salt: Uint8Array;
}

export interface EncryptedPayload {
  iv: string;         // base64
  ciphertext: string; // base64
}

/**
 * Helper: Uint8Array -> ArrayBuffer
 */
function u8ToArrayBuffer(u8: Uint8Array): ArrayBuffer {
  const buffer = u8.buffer as ArrayBuffer; // cast away SharedArrayBuffer possibility
  return buffer.slice(u8.byteOffset, u8.byteOffset + u8.byteLength);
}

/**
 * Generate a random salt.
 */
export function generateSalt(length?: number): Uint8Array {
  const size = length ?? 16;
  const salt = new Uint8Array(size);
  window.crypto.getRandomValues(salt);
  return salt;
}

/**
 * Encode string to Uint8Array (UTF-8).
 */
function encode(text: string): Uint8Array {
  return new TextEncoder().encode(text);
}

/**
 * Decode Uint8Array to string (UTF-8).
 */
function decode(data: Uint8Array): string {
  return new TextDecoder().decode(data);
}

/**
 * Convert Uint8Array to base64 string.
 */
function toBase64(bytes: Uint8Array): string {
  // Avoid explicit numeric loop to sidestep TS issues with length typing
  const binary = String.fromCharCode.apply(null, Array.from(bytes) as number[]);
  return window.btoa(binary);
}

/**
 * Convert base64 string to Uint8Array.
 */
function fromBase64(b64: string): Uint8Array {
  const binary = window.atob(b64);
  const bytes = new Uint8Array(binary.length);
  for (let i = 0; i < binary.length; i++) {
    bytes[i] = binary.charCodeAt(i);
  }
  return bytes;
}

/**
 * Derive a key from a master password using PBKDF2.
 * Returns a CryptoKey and the salt used.
 */
export async function deriveKeyFromPassword(
  password: string,
  salt?: Uint8Array
): Promise<DerivedKeyResult> {
  const saltToUse = salt ?? generateSalt();

  const baseKey = await window.crypto.subtle.importKey(
    'raw',
    u8ToArrayBuffer(encode(password)),
    { name: 'PBKDF2' },
    false,
    ['deriveKey']
  );

  const derivedKey = await window.crypto.subtle.deriveKey(
    {
      name: 'PBKDF2',
      salt: u8ToArrayBuffer(saltToUse),
      iterations: PBKDF2_ITERATIONS,
      hash: PBKDF2_HASH,
    },
    baseKey,
    {
      name: AES_ALGO,
      length: KEY_LENGTH_BITS,
    },
    false,
    ['encrypt', 'decrypt']
  );

  return {
    key: derivedKey,
    salt: saltToUse,
  };
}

/**
 * Generate a random AES-GCM vault key.
 */
export async function generateVaultKey(): Promise<CryptoKey> {
  return window.crypto.subtle.generateKey(
    {
      name: AES_ALGO,
      length: KEY_LENGTH_BITS,
    },
    true, // extractable
    ['encrypt', 'decrypt']
  );
}

/**
 * Export a CryptoKey to raw bytes (Uint8Array).
 */
export async function exportKey(key: CryptoKey): Promise<Uint8Array> {
  const raw = await window.crypto.subtle.exportKey('raw', key);
  return new Uint8Array(raw as ArrayBuffer);
}

/**
 * Import a raw key (Uint8Array) as AES-GCM key.
 */
export async function importVaultKey(raw: Uint8Array): Promise<CryptoKey> {
  return window.crypto.subtle.importKey(
    'raw',
    u8ToArrayBuffer(raw),
    { name: AES_ALGO },
    true,
    ['encrypt', 'decrypt']
  );
}

/**
 * Encrypt a text using an AES-GCM CryptoKey.
 */
export async function encryptText(
  key: CryptoKey,
  plaintext: string
): Promise<EncryptedPayload> {
  const iv = new Uint8Array(IV_LENGTH_BYTES);
  window.crypto.getRandomValues(iv);

  const ciphertextBuffer = await window.crypto.subtle.encrypt(
    {
      name: AES_ALGO,
      iv: u8ToArrayBuffer(iv),
    },
    key,
    u8ToArrayBuffer(encode(plaintext))
  );

  const ciphertext = new Uint8Array(ciphertextBuffer as ArrayBuffer);

  return {
    iv: toBase64(iv),
    ciphertext: toBase64(ciphertext),
  };
}

/**
 * Decrypt an EncryptedPayload back to text using AES-GCM CryptoKey.
 */
export async function decryptText(
  key: CryptoKey,
  payload: EncryptedPayload
): Promise<string> {
  const iv = fromBase64(payload.iv);
  const ciphertext = fromBase64(payload.ciphertext);

  const plaintextBuffer = await window.crypto.subtle.decrypt(
    {
      name: AES_ALGO,
      iv: u8ToArrayBuffer(iv),
    },
    key,
    u8ToArrayBuffer(ciphertext)
  );

  return decode(new Uint8Array(plaintextBuffer as ArrayBuffer));
}