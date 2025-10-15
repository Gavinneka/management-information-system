import { useEffect, useState } from "react";
import api from "../services/api";
import { Link } from "react-router-dom";

export default function UserList() {
  const [users, setUsers] = useState([]);

  const getUsers = async () => {
    const res = await api.get("/users");
    setUsers(res.data);
  };

  const deleteUser = async (id) => {
    if (!confirm("Yakin mau hapus user ini?")) return;
    await api.delete(`/users/${id}`);
    getUsers();
  };

  useEffect(() => {
    getUsers();
  }, []);

  return (
    <div className="container">
      <h1>Manajemen User</h1>
      <Link to="/new" className="btn">+ Tambah User</Link>
      <table>
        <thead>
          <tr>
            <th>ID</th><th>Username</th><th>Nama</th><th>Email</th><th>Aksi</th>
          </tr>
        </thead>
        <tbody>
          {users.map(u => (
            <tr key={u.ID}>
              <td>{u.ID}</td>
              <td>{u.username}</td>
              <td>{u.name}</td>
              <td>{u.email}</td>
              <td>
                <Link to={`/edit/${u.ID}`} className="btn">Edit</Link>
                <button onClick={() => deleteUser(u.ID)} className="btn btn-delete">Hapus</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  );
}
