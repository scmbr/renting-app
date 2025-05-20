import { Routes, Route, Navigate } from 'react-router-dom';
import HomePage from '@/pages/Home/HomePage.jsx';

const App = () => (
  <Routes>
    <Route path="/" element={<Navigate to="/moskva" replace />} />
    <Route path="/:citySlug" element={<HomePage />} />
  </Routes>
);

export default App;
