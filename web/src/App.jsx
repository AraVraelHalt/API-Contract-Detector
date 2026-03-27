import { useEffect, useState } from 'react';

function App() {
  const [changes, setChanges] = useState([]);

  useEffect(() => {
    fetch("http://localhost:8080/changes")
      .then(res => res.json())
      .then(data => setChanges(data || []));
  }, []);

  return (
    <div style={{ padding: "20px" }}>
      <h1>API Contract Break Detector</h1>

      {changes.length === 0 ? (
        <p>No breaking changes detected</p>
      ) : (
        <ul>
          {changes.map((c, index) => (
            <li key={index}>
              <strong>{c.endpoint}</strong>: {c.change}
              <br />
              <small>{c.created_at}</small>
            </li>
          ))}
        </ul>
      )}
    </div>
  );
}

export default App;
