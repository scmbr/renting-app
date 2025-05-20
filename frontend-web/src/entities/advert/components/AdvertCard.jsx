const AdvertCard = ({ advert }) => {
  const { title, rent, apartment } = advert;

  return (
    <div className="border rounded-xl p-4 shadow-md">
      <h3 className="text-lg font-bold">{title}</h3>
      <p>{apartment.city}, {apartment.street} {apartment.house}</p>
      <p>Этаж: {apartment.floor} | Комнат: {apartment.rooms}</p>
      <p className="font-semibold">{rent.toLocaleString()} ₽/мес</p>
    </div>
  );
};

export default AdvertCard;
