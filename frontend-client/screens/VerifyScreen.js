import React, { useState } from 'react';
import { View, Text, TextInput, Button, Alert } from 'react-native';
import axios from 'axios';

const apiUrl = process.env.EXPO_PUBLIC_API_URL;

export default function VerifyScreen() {
  const [code, setCode] = useState('');

  const handleVerify = async () => {
    try {
      await axios.post(`${apiUrl}/auth/verify`, { code });
      Alert.alert('Успех', 'Вы успешно зарегистрированы и вошли в систему!');
    } catch (error) {
      Alert.alert('Ошибка верификации', error.response?.data?.message || 'Неверный код');
    }
  };

  return (
    <View style={{ padding: 20 }}>
      <Text>Введите код верификации</Text>
      <TextInput value={code} onChangeText={setCode} keyboardType="numeric" />
      <Button title="Подтвердить" onPress={handleVerify} />
    </View>
  );
}
