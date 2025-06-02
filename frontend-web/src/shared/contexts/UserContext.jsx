import { createContext, useContext, useState, useEffect } from "react";

const UserContext = createContext(null);

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
  const storedName = localStorage.getItem("name");
  const storedSurname = localStorage.getItem("surname");
  const storedAvatarUrl = localStorage.getItem("avatarUrl");
  if (storedName) {
    setUser({
      name: storedName,
      surname: storedSurname,
      avatarUrl: storedAvatarUrl,
    });
  }
}, []);

  const login = (name, surname, avatarUrl) => {
  localStorage.setItem("name", name);
  localStorage.setItem("surname", surname);
  localStorage.setItem("avatarUrl", avatarUrl);
  setUser({ name, surname, avatarUrl });
};

  const logout = () => {
    localStorage.removeItem("name");
    localStorage.removeItem("surname");
    localStorage.removeItem("accessToken");
    setUser(null, null);
  };

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
