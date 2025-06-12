import { createContext, useContext, useState, useEffect } from "react";

const UserContext = createContext(null);

export const UserProvider = ({ children }) => {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const storedName = localStorage.getItem("name");
    const storedSurname = localStorage.getItem("surname");
    const storedAvatarUrl = localStorage.getItem("avatarUrl");
    const storedCity = localStorage.getItem("city");
    if (storedName && storedSurname && storedCity) {
      const avatar =
        storedAvatarUrl && storedAvatarUrl !== ""
          ? storedAvatarUrl
          : "https://storage.yandexcloud.net/profile-pictures/user.png";

      setUser({
        name: storedName,
        surname: storedSurname,
        avatarUrl: avatar,
        city: storedCity,
      });
    }
  }, []);

  const login = (name, surname, avatarUrl, city) => {
    const avatar =
      avatarUrl || "https://storage.yandexcloud.net/profile-pictures/user.png";
    localStorage.setItem("name", name);
    localStorage.setItem("surname", surname);
    localStorage.setItem("avatarUrl", avatar);
    localStorage.setItem("city", city);
    setUser({ name, surname, avatarUrl: avatar, city });
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
