export type User = {
  id: number;
  name: string;
  roll: number;
  age: number;
  dept: string;
};

const API_BASE = "http://localhost:8080";

export async function getProducts(): Promise<User[]> {
  const res = await fetch(`${API_BASE}/`, {
    method: "GET",
    cache: "no-store",
  });
  if (!res.ok) {
    throw new Error(`GET / failed: ${res.status}`);
  }
  return res.json();
}

export async function createProduct(
  payload: Omit<User, "id">
): Promise<User> {
  const res = await fetch(`${API_BASE}/product`, {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify(payload),
  });
  if (!res.ok) {
    const text = await res.text();
    throw new Error(`POST /product failed: ${res.status} ${text}`);
  }
  return res.json();
}