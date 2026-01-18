// src/api/auth.ts

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export interface SignupPayload {
  email: string;
  password: string; // login password
  encryptedVaultKey: string;
  salt: string; // base64 salt for master key derivation
}

export interface SignupResponse {
  userId: number;
  email: string;
}

export interface LoginPayload {
  email: string;
  password: string; // login password
}

export interface LoginResponse {
  userId: number;
  email: string;
  token: string;
  encryptedVaultKey: string;
  salt: string;
}

export async function signup(payload: SignupPayload): Promise<SignupResponse> {
  const res = await fetch(`${API_BASE_URL}/auth/signup`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    const data = await res.json().catch(() => null);
    throw new Error(data?.error || 'Signup failed');
  }

  return res.json();
}

export async function login(payload: LoginPayload): Promise<LoginResponse> {
  const res = await fetch(`${API_BASE_URL}/auth/login`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    const data = await res.json().catch(() => null);
    throw new Error(data?.error || 'Login failed');
  }

  return res.json();
}