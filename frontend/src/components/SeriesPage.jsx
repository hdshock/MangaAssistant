import React, { useEffect, useState } from "react";
import { Link } from "react-router-dom";

const SeriesPage = () => {
  const [series, setSeries] = useState([]);

  useEffect(() => {
    // Fetch the list of series from the backend
    fetch("http://localhost:2828/api/series")
      .then((response) => response.json())
      .then((data) => setSeries(data))
      .catch((error) => console.error("Failed to fetch series:", error));
  }, []);

  return (
    <div className="series-page">
      <h1>Series</h1>
      <div className="series-grid">
        {series.map((s) => (
          <Link key={s.Series} to={`/series/${encodeURIComponent(s.Series)}`} className="series-card">
            <img src={s.Cover} alt={s.Series} className="series-poster" />
            <h2>{s.Series}</h2>
            <p>{s.Publisher}</p>
          </Link>
        ))}
      </div>
    </div>
  );
};

export default SeriesPage; 