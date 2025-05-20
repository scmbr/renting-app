import React from 'react';
import { Routes, Route } from 'react-router-dom';
import HomePage from '../pages/Home';

const Router = () => (
  <Routes>
    <Route path="/" element={<HomePage />} />
  </Routes>
);

export default Router;
