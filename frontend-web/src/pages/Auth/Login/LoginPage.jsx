import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import styles from "./LoginPage.module.css";
import api from "@/shared/api/axios";
const LoginPage = () => {
  const navigate = useNavigate();
  const [email, setEmail] = useState("");
  const [password, setPassword] = useState("");
  const [error, setError] = useState(null);

  const handleLogin = async (e) => {
    e.preventDefault();
    setError(null);

    try {
      const res = await api.post("/auth/sign-in", { email, password });
      localStorage.setItem("accessToken", res.data.accessToken);
      navigate("/");
    } catch (err) {
      console.error(err);
      setError("Неверный email или пароль");
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.left}>
        <div className={styles.container}>
          <img
            src="/images/ROOMY.png"
            alt="Login Icon"
            className={styles.icon}
          />
          {/* <h2 className={styles.title}>Вход</h2> */}

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
