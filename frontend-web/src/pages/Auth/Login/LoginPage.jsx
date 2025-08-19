import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import styles from "./LoginPage.module.css";
import api from "@/shared/api/axios";
import { useUser } from "@/shared/contexts/UserContext";
import { useCityStore } from "@/stores/useCityStore";
import { nameToSlug } from "@/shared/constants/cities";
const LoginPage = () => {
  const navigate = useNavigate();
  const { login } = useUser();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);
  const city = useCityStore((state) => state.city);
  const citySlug = nameToSlug(city || "Москва");
  const handleLogin = async (e) => {
    e.preventDefault();
    setError(null);

    try {
      const res = await api.post("/auth/sign-in", { email, password });
      const token = res.data.accessToken;
      localStorage.setItem("accessToken", token);

      const userRes = await api.get("/me", {
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      const name = userRes.data.name;
      const surname = userRes.data.surname;
      const avatarUrl = userRes.data.profile_picture;
      const city = userRes.data.city;

      const avatarUrlToUse = avatarUrl
        ? avatarUrl
        : "/images/no-photo.png";

      login(name, surname, avatarUrlToUse, city);
      navigate(`/${city}`);
    } catch (err) {
      console.error(err);
      setError("Неверный email или пароль");
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.left}>
        <div className={styles.container}>
          <Link to={`/${citySlug}`}>
            <img
              src="/images/logo.png"
              alt="Login Icon"
              className={styles.icon}
            />
          </Link>

          {error && <p className={styles.error}>{error}</p>}

          <form onSubmit={handleLogin} className={styles.form}>
            <input
              type="email"
              placeholder="Email"
              className={styles.input}
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              required
            />
            <input
              type="password"
              placeholder="Пароль"
              className={styles.input}
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              required
            />
            <button type="submit" className={styles.button}>
              Войти
            </button>
          </form>

          <p className={styles.registerLink}>
            Или{" "}
            <Link to="/register" className={styles.link}>
              зарегистрируйтесь
            </Link>
          </p>
        </div>
      </div>
      <div className={styles.right}></div>
    </div>
  );
};

export default LoginPage;
