import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import styles from "./RegisterPage.module.css";
import api from "@/shared/api/axios";

const RegisterPage = () => {
  const navigate = useNavigate();
  const [step, setStep] = useState(1);
  const [form, setForm] = useState({
    name: "",
    surname: "",
    birthdate: "",
    email: "",
    password: "",
    confirm: "",
  });
  const [error, setError] = useState(null);

  const handleChange = (e) => {
    const { name, value } = e.target;
    setForm((prev) => ({ ...prev, [name]: value }));
  };

  const handleNext = () => {
    if (!form.name || !form.surname) {
      setError("Имя и фамилия обязательны");
      return;
    }
    setError(null);
    setStep(2);
  };

  const handleRegister = async (e) => {
    e.preventDefault();
    setError(null);

    if (form.password !== form.confirm) {
      setError("Пароли не совпадают");
      return;
    }
    const normalizedBirthdate = form.birthdate
      ? new Date(form.birthdate).toISOString()
      : undefined;
    try {
      await api.post("/auth/sign-up", {
        name: form.name,
        surname: form.surname,
        email: form.email,
        password: form.password,
        birthdate: normalizedBirthdate || undefined,
      });
      navigate("/verify");
    } catch (err) {
      console.error(err);
      setError("Ошибка при регистрации");
    }
  };

  return (
    <div className={styles.wrapper}>
      <div className={styles.left}>
        <div className={styles.container}>
          <Link to="/">
            <img
              src="/images/logo.png"
              alt="Login Icon"
              className={styles.logo}
            />
          </Link>

          <div className={styles.stepHeader}>
            {step === 2 && (
              <button className={styles.backButton} onClick={() => setStep(1)}>
                <img
                  src="/icons/back.svg"
                  alt="Назад"
                  className={styles.icon}
                />
              </button>
            )}
            <div className={styles.stepIndicators}>
              <div
                className={`${styles.dot} ${step === 1 ? styles.active : ""}`}
              />
              <div
                className={`${styles.dot} ${step === 2 ? styles.active : ""}`}
              />
            </div>
          </div>

          {error && <p className={styles.error}>{error}</p>}

          <form onSubmit={handleRegister} className={styles.form}>
            {step === 1 && (
              <>
                <input
                  type="text"
                  name="name"
                  placeholder="Имя"
                  className={styles.input}
                  value={form.name}
                  onChange={handleChange}
                  required
                />
                <input
                  type="text"
                  name="surname"
                  placeholder="Фамилия"
                  className={styles.input}
                  value={form.surname}
                  onChange={handleChange}
                  required
                />

                <input
                  type="date"
                  name="birthdate"
                  className={styles.input}
                  value={form.birthdate}
                  onChange={handleChange}
                />
                <button
                  type="button"
                  className={styles.button}
                  onClick={handleNext}
                >
                  Далее
                </button>
              </>
            )}

            {step === 2 && (
              <>
                <input
                  type="email"
                  name="email"
                  placeholder="Email"
                  className={styles.input}
                  value={form.email}
                  onChange={handleChange}
                  required
                />
                <input
                  type="password"
                  name="password"
                  placeholder="Пароль"
                  className={styles.input}
                  value={form.password}
                  onChange={handleChange}
                  required
                />
                <input
                  type="password"
                  name="confirm"
                  placeholder="Повторите пароль"
                  className={styles.input}
                  value={form.confirm}
                  onChange={handleChange}
                  required
                />
                <button type="submit" className={styles.button}>
                  Зарегистрироваться
                </button>
              </>
            )}
          </form>

          <p className={styles.registerLink}>
            Уже есть аккаунт?{" "}
            <Link to="/login" className={styles.link}>
              Войти
            </Link>
          </p>
        </div>
      </div>
      <div className={styles.right}></div>
    </div>
  );
};

export default RegisterPage;
