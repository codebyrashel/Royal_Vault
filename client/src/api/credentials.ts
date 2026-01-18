// src/api/credentials.ts

const API_BASE_URL = import.meta.env.VITE_API_BASE_URL || 'http://localhost:8080';

export interface CredentialResponse {
  id: number;
  title: string;
  url?: string | null;
  encryptedUsername: string;
  encryptedPassword: string;
  encryptedNotes?: string | null;
  folderId?: number | null;
}

export interface CreateCredentialPayload {
  title: string;
  url?: string | null;
  encryptedUsername: string;
  encryptedPassword: string;
  encryptedNotes?: string | null;
  folderId?: number | null;
}

export async function fetchCredentials(token: string): Promise<CredentialResponse[]> {
  const res = await fetch(`${API_BASE_URL}/credentials`, {
    method: 'GET',
    headers: {
      Authorization: `Bearer ${token}`,
    },
  });

  if (!res.ok) {
    const data = await res.json().catch(() => null);
    throw new Error(data?.error || 'Failed to fetch credentials');
  }

  return res.json();
}

export async function createCredential(
  token: string,
  payload: CreateCredentialPayload
): Promise<CredentialResponse> {
  const res = await fetch(`${API_BASE_URL}/credentials`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      Authorization: `Bearer ${token}`,
    },
    body: JSON.stringify(payload),
  });

  if (!res.ok) {
    const data = await res.json().catch(() => null);
    throw new Error(data?.error || 'Failed to create credential');
  }

  return res.json();
}