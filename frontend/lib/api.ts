const BACKEND_URL = process.env.BACKEND_URL ?? "http://localhost:8080";

type FetchOptions = {
  method?: string;
  body?: unknown;
  token?: string;
};

export async function backendFetch<T>(
  path: string,
  { method = "GET", body, token }: FetchOptions = {}
): Promise<T> {
  const res = await fetch(`${BACKEND_URL}${path}`, {
    method,
    headers: {
      "Content-Type": "application/json",
      ...(token ? { Authorization: `Bearer ${token}` } : {}),
    },
    ...(body ? { body: JSON.stringify(body) } : {}),
  });

  if (!res.ok) {
    const error = await res.text();
    throw new Error(error || `Backend error: ${res.status}`);
  }

  // 204 No Content (e.g. logout) — return empty
  if (res.status === 204) return {} as T;

  return res.json() as Promise<T>;
}