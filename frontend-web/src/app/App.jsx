import { Routes, Route, Navigate } from "react-router-dom";
import HomePage from "@/pages/Home/HomePage.jsx";
import RegisterPage from "@/pages/Auth/Register/RegisterPage.jsx";
import LoginPage from "@/pages/Auth/Login/LoginPage.jsx";
import AdvertPage from "@/pages/Advert/AdvertPage";
import AddAdvertPage from "@/pages/AddAdvert/AddAdvertPage";
import AddApartmentPage from "@/pages/AddApartment/AddApartmentPage";
import MyAdvertsPage from "@/pages/MyAdvertsPage/MyAdvertsPage";
import MyApartmentsPage from "@/pages/MyApartmentsPage/MyApartmentsPage";
import FavoritesPage from "@/pages/FavoritesPage/FavoritesPage";
import SettingsPage from "@/pages/SettingsPage/SettingsPage";
import NotificationsPage from "@/pages/NotificationsPage/NotificationsPage";
import "@/index.css";
import { UserProvider } from "@/shared/contexts/UserContext";
import VerifyPage from "@/pages/Verify/VerifyPage";
import EditAdvertPage from "@/pages/EditAdvert/EditAdvertPage";
import EditApartmentPage from "@/pages/EditApartment/EditApartmentPage";
const App = () => (
  <Routes>
    <Route path="/" element={<Navigate to="/moskva" replace />} />
    <Route path="/my/apartment/add" element={<AddApartmentPage />} />
    <Route path="/my/apartment/edit/:id" element={<EditApartmentPage />} />
    <Route path="/my/advert/add" element={<AddAdvertPage />} />
    <Route path="/my/advert/edit/:id" element={<EditAdvertPage />} />
    <Route path="/advert/:id" element={<AdvertPage />} />

    <Route path="/favorites" element={<FavoritesPage />} />
    <Route path="/settings" element={<SettingsPage />} />

    <Route path="/my/advert" element={<MyAdvertsPage />} />
    <Route path="/my/apartment" element={<MyApartmentsPage />} />
    <Route path="/verify" element={<VerifyPage />} />
    <Route path="/:citySlug" element={<HomePage />} />
    <Route path="/register" element={<RegisterPage />} />
    <Route path="/login" element={<LoginPage />} />
    <Route path="/notifications" element={<NotificationsPage />} />
    {/* <Route path="/forgot-password" element={<ForgotPasswordPage />} />
    <Route path="/reset-password" element={<ResetPasswordPage />} /> */}
  </Routes>
);

export default App;
