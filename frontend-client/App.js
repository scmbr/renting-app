import React, { useState } from 'react';
import { View, Text, TextInput, Button, Alert } from 'react-native';
import axios from 'axios';

const apiUrl = process.env.EXPO_PUBLIC_API_URL;
export default function App() {
  const [screen, setScreen] = useState('register');
  const [userData, setUserData] = useState({
    name: '',
    surname: '',
    email: '',
    password: '',
    birthdate: '',
  });
  const [code, setCode] = useState('');

  const handleRegister = async () => {
    try {
      console.log('userData', userData);
      console.log('ISO Birthdate:', new Date(userData.birthdate).toISOString());
      console.log(`${apiUrl}/auth/sign-up`);
      await axios.post(`${apiUrl}/auth/sign-up`, {
        ...userData,
        birthdate: new Date(userData.birthdate).toISOString(),
      });
  
      setScreen('verify');
    } catch (error) {
      console.error('Registration error:', error.response?.data || error.message || error);
      Alert.alert(
        'Ошибка регистрации',
        error.response?.data?.message || error.message || 'Что-то пошло не так'
      );
    }
  };

  const handleVerify = async () => {
    try {
      await axios.post(`${apiUrl}/auth/verify`, { code });
      Alert.alert('Успех', 'Вы успешно зарегистрированы и вошли в систему!');
    } catch (error) {
      Alert.alert('Ошибка верификации', error.response?.data?.message || 'Неверный код');
    }
  };

  if (screen === 'register') {
    return (
      <View style={{ padding: 20 }}>
        <Text>Имя</Text>
        <TextInput value={userData.name} onChangeText={(text) => setUserData({ ...userData, name: text })} />

        <Text>Фамилия</Text>
        <TextInput value={userData.surname} onChangeText={(text) => setUserData({ ...userData, surname: text })} />

        <Text>Email</Text>
        <TextInput value={userData.email} onChangeText={(text) => setUserData({ ...userData, email: text })} keyboardType="email-address" />

        <Text>Пароль</Text>
        <TextInput value={userData.password} onChangeText={(text) => setUserData({ ...userData, password: text })} secureTextEntry />

        <Text>Дата рождения (ГГГГ-ММ-ДД)</Text>
        <TextInput value={userData.birthdate} onChangeText={(text) => setUserData({ ...userData, birthdate: text })} />

        <Button title="Зарегистрироваться" onPress={handleRegister} />
      </View>
    );
  }

  return (
    <View style={{ padding: 20 }}>
      <Text>Введите код верификации</Text>
      <TextInput value={code} onChangeText={setCode} keyboardType="numeric" />
      <Button title="Подтвердить" onPress={handleVerify} />
    </View>
  );
}
