import React, { useEffect, useState } from "react";
import { useParams } from "react-router-dom";

const SeriesDetail = () => {
  const { seriesName } = useParams();
  const [series, setSeries] = useState(null);

  useEffect(() => {
    // Fetch the series detail from the backend
    fetch(`http://localhost:2828/api/series/detail?name=${encodeURIComponent(seriesName)}`)
      .then((response) => response.json())
      .then((data) => setSeries(data))
      .catch((error) => console.error("Failed to fetch series detail:", error));
  }, [seriesName]);

  if (!series) {
    return <div>Loading...</div>;
  }

  return (
    <div className="series-detail">
      <h1>{series.Series}</h1>
      <div className="series-info">
        <img src={series.Cover} alt={series.Series} className="series-poster" />
        <div className="series-metadata">
          <p><strong>Publisher:</strong> {series.Publisher}</p>
          <p><strong>Volumes:</strong> {series.Volumes}</p>
          <p><strong>Chapters:</strong> {series.Chapters}</p>
        </div>
      </div>
    </div>
  );
};

export default SeriesDetail; 