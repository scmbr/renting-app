import { Routes, Route, Navigate } from "react-router-dom";
import HomePage from "@/pages/Home/HomePage.jsx";
import RegisterPage from "@/pages/Auth/Register/RegisterPage.jsx";
import LoginPage from "@/pages/Auth/Login/LoginPage.jsx";
import AdvertPage from "@/pages/Advert/AdvertPage";
import AddAdvertPage from "@/pages/AddAdvert/AddAdvertPage";
import AddApartmentPage from "@/pages/Addapartment/AddapartmentPage";
import MyAdvertsPage from "@/pages/MyAdvertsPage/MyAdvertsPage";
import "@/index.css";
import { UserProvider } from "@/shared/contexts/UserContext";
import VerifyPage from "@/pages/Verify/VerifyPage";
const App = () => (
  <Routes>
    <Route path="/" element={<Navigate to="/moskva" replace />} />
    <Route path="/advert/:id" element={<AdvertPage />} />
    <Route path="/advert/add" element={<AddAdvertPage />} />
    <Route path="/apartment/add" element={<AddApartmentPage />} />
    <Route path="/my/advert" element={<MyAdvertsPage />} />
    <Route path="/verify" element={<VerifyPage />} />
    <Route path="/:citySlug" element={<HomePage />} />
    <Route path="/register" element={<RegisterPage />} />
    <Route path="/login" element={<LoginPage />} />

    {/* <Route path="/forgot-password" element={<ForgotPasswordPage />} />
    <Route path="/reset-password" element={<ResetPasswordPage />} /> */}
  </Routes>
);

export default App;
