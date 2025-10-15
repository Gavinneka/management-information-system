import { useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";
import api from "../services/api";

export default function UserForm() {
  const { id } = useParams();
  const navigate = useNavigate();
  const [user, setUser] = useState({ username: "", name: "", email: "", password: "" });

  useEffect(() => {
    if (id) {
      api.get(`/users/${id}`).then(res => setUser(res.data));
    }
  }, [id]);

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (id) await api.put(`/users/${id}`, user);
    else await api.post("/users", user);
    navigate("/");
  };

  return (
    <div className="container">
      <h2>{id ? "Edit User" : "Tambah User"}</h2>
      <form onSubmit={handleSubmit}>
        <input value={user.username} onChange={e => setUser({...user, username: e.target.value})} placeholder="Username" required />
        <input value={user.name} onChange={e => setUser({...user, name: e.target.value})} placeholder="Nama" />
        <input type="email" value={user.email} onChange={e => setUser({...user, email: e.target.value})} placeholder="Email" required />
        <input type="password" value={user.password} onChange={e => setUser({...user, password: e.target.value})} placeholder="Password" />
        <button type="submit">Simpan</button>
      </form>
    </div>
  );
}
