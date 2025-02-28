import React from 'react';
import { Link } from 'react-router-dom';

const Sidebar = () => {
  return (
    <aside className="w-64 bg-gray-800 text-white p-4">
      <h2 className="text-lg font-semibold">Library</h2>
      <ul className="mt-4 space-y-2">
        <li>
          <Link to="/" className="block p-2 hover:bg-gray-700 rounded">Home</Link>
        </li>
        <li>
          <Link to="/series" className="block p-2 hover:bg-gray-700 rounded">Series</Link>
        </li>
        <li>
          <Link to="/collections" className="block p-2 hover:bg-gray-700 rounded">Collections</Link>
        </li>
        <li>
          <Link to="/settings" className="block p-2 hover:bg-gray-700 rounded">Settings</Link>
        </li>
      </ul>
    </aside>
  );
};

export default Sidebar; 