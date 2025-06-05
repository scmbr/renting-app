import React, { useState } from "react";
import { useNavigate } from "react-router-dom";
import styles from "./VerifyPage.module.css";
import api from "@/shared/api/axios";
import { useUser } from "@/shared/contexts/UserContext";

const VerifyPage = () => {
  const [code, setCode] = useState("");
  const [error, setError] = useState(null);
  const navigate = useNavigate();
  const { login } = useUser();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError(null);

    try {
      const res = await api.post("/auth/verify", { code });
      const token = res.data.accessToken;

      localStorage.setItem("accessToken", token);
      api.defaults.headers.common["Authorization"] = `Bearer ${token}`;

      const userRes = await api.get("/me");
      const { name, surname, profile_picture,city } = userRes.data;

      login(name, surname, profile_picture,city);
      navigate("/");
    } catch (err) {
      console.error(err);
      setError("Неверный код верификации");
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.left}>
        <div className={styles.container}>
          <img src="/images/logo.png" alt="Logo" className={styles.logo} />

          <h2 className={styles.title}>Подтверждение аккаунта</h2>

          {error && <p className={styles.error}>{error}</p>}

          <form onSubmit={handleSubmit} className={styles.form}>
            <input
              type="text"
              placeholder="Код верификации"
              className={styles.input}
              value={code}
              onChange={(e) => setCode(e.target.value)}
              required
            />
            <button type="submit" className={styles.button}>
              Подтвердить
            </button>
          </form>
        </div>
      </div>
      <div className={styles.right}></div>
    </div>
  );
};

export default VerifyPage;
