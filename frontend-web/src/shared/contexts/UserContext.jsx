import { createContext, useContext, useState, useEffect } from "react";

const UserContext = createContext(null);

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const storedName = localStorage.getItem("name");
    const storedSurname = localStorage.getItem("surname");
    const storedAvatarUrl = localStorage.getItem("avatarUrl");
    const storedCity = localStorage.getItem("city");
    if (storedName && storedSurname && storedAvatarUrl && storedCity) {
      setUser({
        name: storedName,
        surname: storedSurname,
        avatarUrl: storedAvatarUrl,
        city: storedCity,
      });
    }
  }, []);

  const login = (name, surname, avatarUrl, city) => {
    localStorage.setItem("name", name);
    localStorage.setItem("surname", surname);
    localStorage.setItem("avatarUrl", avatarUrl);
    localStorage.setItem("city", city);
    setUser({ name, surname, avatarUrl, city });
  };

  const logout = () => {
    localStorage.removeItem("name");
    localStorage.removeItem("surname");
    localStorage.removeItem("accessToken");
    localStorage.removeItem("city");
    setUser(null, null, null, null);
  };

  return (
    <UserContext.Provider value={{ user, login, logout }}>
      {children}
    </UserContext.Provider>
  );
};

export const useUser = () => useContext(UserContext);
