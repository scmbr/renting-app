import React, { useState } from 'react';
import { TouchableOpacity, Text, View, TextInput, Alert, StyleSheet, ScrollView,Image } from 'react-native';
import axios from 'axios';
import AppLoading from 'expo-app-loading'; // Для загрузки шрифта до рендера компонента
import { useFonts } from 'expo-font';
const apiUrl = process.env.EXPO_PUBLIC_API_URL;

export default function RegisterScreen({ navigation }) {
  const [step, setStep] = useState(1); // Шаг формы, 1 - первый экран, 2 - второй экран
  const [userData, setUserData] = useState({
    name: '',
    surname: '',
    email: '',
    password: '',
    birthdate: '',
  });
  const [fontsLoaded] = useFonts({
    'TikTokSans': require('../assets/fonts/TikTokSans-VF-v3.3.ttf'), // Загружаем шрифт
    'Futura': require('../assets/fonts/zhuan_FuturaNowScript-XBd.ttf'),
  });
 
  if (!fontsLoaded) {
    return <AppLoading />;
  }
  const handleRegisterStep1 = () => {
    // Проверка на наличие email и пароля
    if (!userData.email || !userData.password) {
      Alert.alert('Ошибка', 'Пожалуйста, заполните все поля');
      return;
    }
    setStep(2); // Переход ко второй части формы
  };

  const handleRegisterStep2 = async () => {
    // Проверка на наличие имени, фамилии и даты рождения
    if (!userData.name || !userData.surname || !userData.birthdate) {
      Alert.alert('Ошибка', 'Пожалуйста, заполните все поля');
      return;
    }

    try {
      const dataToSend = { ...userData };

      if (userData.birthdate.trim() !== '') {
        const [day, month, year] = userData.birthdate.split('.');
        dataToSend.birthdate = `${year}-${month}-${day}T00:00:00Z`;
      } else {
        delete dataToSend.birthdate;
      }

      await axios.post(`${apiUrl}/auth/sign-up`, dataToSend);
      navigation.navigate('Verify');
    } catch (error) {
      Alert.alert('Ошибка регистрации', error.response?.data?.message || 'Что-то пошло не так');
    }
  };

  const formatDate = (text) => {
    let formattedDate = text.replace(/\D/g, '');
    if (formattedDate.length > 4) {
      formattedDate = formattedDate.replace(/(\d{2})(\d{2})(\d{4})/, '$1.$2.$3');
    } else if (formattedDate.length > 2) {
      formattedDate = formattedDate.replace(/(\d{2})(\d{2})/, '$1.$2');
    }
    return formattedDate;
  };

  return (
    <ScrollView contentContainerStyle={styles.container}>
         <Image source={require('../assets/ROOMY.png')} style={styles.logo} />
      <View style={styles.form}>
        {step === 1 ? (
          // Первая часть формы (email и пароль)
          <>
            <TextInput
              style={styles.input}
              placeholder="Email"
              placeholderTextColor="#626262"
              value={userData.email}
              onChangeText={(text) => setUserData({ ...userData, email: text })}
              keyboardType="email-address"
            />
            <TextInput
              style={styles.input}
              placeholder="Пароль"
              placeholderTextColor="#626262"
              value={userData.password}
              onChangeText={(text) => setUserData({ ...userData, password: text })}
              secureTextEntry
            />
            <TouchableOpacity style={styles.button} onPress={handleRegisterStep1}>
              <Text style={styles.buttonText}>Далее</Text>
            </TouchableOpacity>
          </>
        ) : (
          // Вторая часть формы (имя, фамилия и дата рождения)
          <>
            <TextInput
              style={styles.input}
              placeholder="Имя"
              placeholderTextColor="#626262"
              value={userData.name}
              onChangeText={(text) => setUserData({ ...userData, name: text })}
            />
            <TextInput
              style={styles.input}
              placeholder="Фамилия"
              placeholderTextColor="#626262"
              value={userData.surname}
              onChangeText={(text) => setUserData({ ...userData, surname: text })}
            />
            <TextInput
              style={styles.input}
              placeholder="Дата рождения (ДД.ММ.ГГГГ)"
              placeholderTextColor="#626262"
              value={userData.birthdate}
              onChangeText={(text) => setUserData({ ...userData, birthdate: formatDate(text) })}
              keyboardType="numeric"
            />
            <TouchableOpacity style={styles.button} onPress={handleRegisterStep2}>
              <Text style={styles.buttonText}>Зарегистрироваться</Text>
            </TouchableOpacity>
          </>
        )}
      </View>
    </ScrollView>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
    padding: 20,
    backgroundColor: '#ffffff',
  },
  form: {
    width: '100%',
    maxWidth: 400,
    padding: 20,
  },
  input: {
    marginTop: 10,
    height: 50,
    borderColor: '#ddd',
    borderWidth: 0,
    borderRadius: 5,
    marginBottom: 16,
    paddingLeft:25,
    paddingHorizontal: 10,
    fontSize: 16,
    fontFamily: 'TikTokSans',
    backgroundColor: '#F1F4FF',
    color: '#000000',
  },
  button: {
    backgroundColor: '#1F41BB',
    paddingVertical: 20,
    borderRadius: 8,
    alignItems: 'center',
    marginTop: 20,
  },
  buttonText: {
    color: '#fff',
    fontSize: 16,
    fontFamily: 'TikTokSans',
  },
  logo: {
    width: 250,
    height: 100,
    marginBottom: 0,
    resizeMode: 'contain',
  },
});
