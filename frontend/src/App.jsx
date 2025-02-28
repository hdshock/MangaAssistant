import React, { useState, useEffect } from 'react';
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';
import Navbar from './components/Navbar';
import Sidebar from './components/Sidebar';
import MangaList from './components/MangaList';
import LibraryManager from './components/LibraryManager';
import CheckDuplicates from './components/CheckDuplicates';
import SeriesPage from './components/SeriesPage';
import SeriesDetail from './components/SeriesDetail';
import Collections from './components/Collections';
import Settings from './components/Settings';

function App() {
  const [mangaList, setMangaList] = useState([]);

  useEffect(() => {
    fetch('http://localhost:2828/api/manga')
      .then((response) => response.json())
      .then((data) => setMangaList(data))
      .catch((error) => console.error('Error fetching manga list:', error));
  }, []);

  return (
    <Router>
      <div className="flex h-screen bg-gray-100">
        <Sidebar />
        <div className="flex-1 flex flex-col overflow-hidden">
          <Navbar />
          <main className="flex-1 overflow-x-hidden overflow-y-auto p-4">
            <Routes>
              <Route path="/" element={<MangaList mangaList={mangaList} />} />
              <Route path="/library-manager" element={<LibraryManager />} />
              <Route path="/check-duplicates" element={<CheckDuplicates />} />
              <Route path="/series" element={<SeriesPage />} />
              <Route path="/series/:seriesName" element={<SeriesDetail />} />
              <Route path="/collections" element={<Collections />} />
              <Route path="/settings" element={<Settings />} />
            </Routes>
          </main>
        </div>
      </div>
    </Router>
  );
}

export default App; 