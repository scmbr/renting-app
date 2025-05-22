import { Routes, Route, Navigate } from "react-router-dom";
import HomePage from "@/pages/Home/HomePage.jsx";
import RegisterPage from "@/pages/Auth/RegisterPage.jsx";
import LoginPage from "@/pages/Auth/Login/LoginPage.jsx";
import AdvertPage from "@/pages/Advert/AdvertPage";
import "@/index.css";
const App = () => (
  <Routes>
    <Route path="/" element={<Navigate to="/moskva" replace />} />
    <Route path="/advert/:id" element={<AdvertPage />} />
    <Route path="/:citySlug" element={<HomePage />} />
    <Route path="/register" element={<RegisterPage />} />
    <Route path="/login" element={<LoginPage />} />

    {/* <Route path="/forgot-password" element={<ForgotPasswordPage />} />
    <Route path="/reset-password" element={<ResetPasswordPage />} /> */}
  </Routes>
);

export default App;
