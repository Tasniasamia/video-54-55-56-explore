"use client";

import { useEffect, useState } from "react";
import { createProduct, getProducts, type User } from "@/lib/api";

export default function HomePage() {
  const [users, setUsers] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [submitting, setSubmitting] = useState(false);
  const [error, setError] = useState<string | null>(null);

  const [form, setForm] = useState({
    name: "",
    roll: "",
    age: "",
    dept: "",
  });

  async function loadUsers() {
    setLoading(true);
    setError(null);
    try {
      const data = await getProducts();
      setUsers(data);
    } catch (e) {
      setError(e instanceof Error ? e.message : "Failed to load");
    } finally {
      setLoading(false);
    }
  }

  useEffect(() => {
    loadUsers();
  }, []);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setSubmitting(true);
    setError(null);
    try {
      await createProduct({
        name: form.name,
        roll: Number(form.roll),
        age: Number(form.age),
        dept: form.dept,
      });
      setForm({ name: "", roll: "", age: "", dept: "" });
      await loadUsers();
    } catch (e) {
      setError(e instanceof Error ? e.message : "Failed to create");
    } finally {
      setSubmitting(false);
    }
  }

  return (
    <main>
      <h1>Ecommerce Admin</h1>
      <p className="subtitle">
        Connected to Go API at <code>:8080</code> via Next.js rewrite
      </p>

      {error && <div className="error">{error}</div>}

      <h2>Create product</h2>
      <form onSubmit={handleSubmit}>
        <label>
          Name
          <input
            required
            value={form.name}
            onChange={(e) => setForm({ ...form, name: e.target.value })}
          />
        </label>
        <label>
          Roll
          <input
            required
            type="number"
            value={form.roll}
            onChange={(e) => setForm({ ...form, roll: e.target.value })}
          />
        </label>
        <label>
          Age
          <input
            required
            type="number"
            value={form.age}
            onChange={(e) => setForm({ ...form, age: e.target.value })}
          />
        </label>
        <label>
          Dept
          <input
            required
            value={form.dept}
            onChange={(e) => setForm({ ...form, dept: e.target.value })}
          />
        </label>
        <button type="submit" disabled={submitting}>
          {submitting ? "Creating..." : "Create"}
        </button>
      </form>

      <div className="toolbar">
        <h2 style={{ margin: 0 }}>Products ({users.length})</h2>
        <button className="refresh" onClick={loadUsers} disabled={loading}>
          {loading ? "Loading..." : "Refresh"}
        </button>
      </div>

      {loading && users.length === 0 ? (
        <p className="subtitle">Loading...</p>
      ) : users.length === 0 ? (
        <p className="subtitle">No products yet.</p>
      ) : (
        users.map((u) => (
          <div className="card" key={u.id}>
            <div className="field">
              <span className="label">ID</span>
              <span className="value">{u.id}</span>
            </div>
            <div className="field">
              <span className="label">Name</span>
              <span className="value">{u.name}</span>
            </div>
            <div className="field">
              <span className="label">Roll</span>
              <span className="value">{u.roll}</span>
            </div>
            <div className="field">
              <span className="label">Age</span>
              <span className="value">{u.age}</span>
            </div>
            <div className="field">
              <span className="label">Dept</span>
              <span className="value">{u.dept}</span>
            </div>
          </div>
        ))
      )}
    </main>
  );
}