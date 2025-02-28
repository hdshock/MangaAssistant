import React, { useState } from 'react';

const MangaList = ({ mangaList }) => {
  const [duplicates, setDuplicates] = useState([]);
  const [scannedManga, setScannedManga] = useState([]);

  const handleCheckDuplicates = () => {
    fetch('/api/manga/check-duplicates')
      .then((response) => response.json())
      .then((data) => setDuplicates(data))
      .catch((error) => console.error('Error checking duplicates:', error));
  };

  const handleScan = async () => {
    try {
      const response = await fetch('http://localhost:2828/scan', {
        method: 'POST',
      });
      if (response.ok) {
        const data = await response.text();
        alert('Scan initiated successfully!');
        // Fetch the scanned manga files
        fetchScannedManga();
      } else {
        alert('Failed to initiate scan.');
      }
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred while initiating the scan.');
    }
  };

  const fetchScannedManga = async () => {
    try {
      const response = await fetch('/api/manga');
      if (response.ok) {
        const data = await response.json();
        setScannedManga(data);
      }
    } catch (error) {
      console.error('Error fetching scanned manga:', error);
    }
  };

  const handleBrowse = async () => {
    try {
      const response = await fetch('http://localhost:2828/browse');
      if (response.ok) {
        const data = await response.text();
        alert(`Files in /mnt/manga: ${data}`);
      } else {
        alert('Failed to browse directory.');
      }
    } catch (error) {
      console.error('Error:', error);
      alert('An error occurred while browsing the directory.');
    }
  };

  return (
    <div className="bg-white p-6 rounded-lg shadow-md">
      <h2 className="text-xl font-semibold mb-4">Manga List</h2>
      <div className="flex space-x-2">
        <button
          onClick={handleCheckDuplicates}
          className="bg-green-500 text-white p-2 rounded hover:bg-green-600"
        >
          Check Duplicates
        </button>
        <button
          onClick={handleScan}
          className="bg-blue-500 text-white p-2 rounded hover:bg-blue-600"
        >
          Scan Library
        </button>
        <button
          onClick={handleBrowse}
          className="bg-purple-500 text-white p-2 rounded hover:bg-purple-600"
        >
          Browse
        </button>
      </div>

      {/* Display scanned manga files */}
      {scannedManga.length > 0 && (
        <div className="mt-4">
          <h3 className="text-lg font-semibold">Scanned Manga:</h3>
          <ul className="space-y-2">
            {scannedManga.map((manga, index) => (
              <li key={index} className="p-2 border rounded">
                <p><strong>Title:</strong> {manga.Title}</p>
                <p><strong>Series:</strong> {manga.Series}</p>
                <p><strong>Volume:</strong> {manga.Volume}</p>
                <p><strong>Chapter:</strong> {manga.Number}</p>
                <p><strong>File:</strong> {manga.FilePath}</p>
              </li>
            ))}
          </ul>
        </div>
      )}

      {/* Display duplicates */}
      {duplicates.length > 0 && (
        <div className="mt-4">
          <h3 className="text-lg font-semibold">Duplicate Manga:</h3>
          <ul className="space-y-2">
            {duplicates.map((manga, index) => (
              <li key={index} className="p-2 border rounded">
                <p><strong>Title:</strong> {manga.Title}</p>
                <p><strong>Series:</strong> {manga.Series}</p>
                <p><strong>Volume:</strong> {manga.Volume}</p>
              </li>
            ))}
          </ul>
        </div>
      )}
    </div>
  );
};

export default MangaList; 