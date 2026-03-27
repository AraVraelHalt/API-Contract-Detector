import { useEffect, useState } from "react";
import { FaCheckCircle, FaExclamationTriangle } from "react-icons/fa";
import "./Dashboard.css";

function Dashboard() {
  const [changes, setChanges] = useState([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState(null);

  useEffect(() => {
    fetch("http://localhost:8080/changes")
      .then((res) => {
        if (!res.ok) throw new Error("Failed to fetch");
        return res.json();
      })
      .then((data) => setChanges(data || []))
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  return (
    <div className="dashboard-container">
      <header>
        <h1>API Contract Break Detector</h1>
      </header>

      {loading ? (
        <div className="status-message">Loading changes...</div>
      ) : error ? (
        <div className="status-message error">Error: {error}</div>
      ) : changes.length === 0 ? (
        <div className="status-message success">
          <FaCheckCircle className="icon" />
          No breaking changes detected
        </div>
      ) : (
        <ul className="changes-list">
          {changes.map((c, index) => (
            <li key={index} className="change-item">
              <div className="change-header">
                <FaExclamationTriangle className="icon warning" />
                <strong>{c.endpoint}</strong>
                <span className="date">{new Date(c.created_at).toLocaleString()}</span>
              </div>
              <div className="change-detail">{c.change}</div>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default Dashboard;
