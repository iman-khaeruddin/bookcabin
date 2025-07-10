import React, { useState } from 'react';
import axios from 'axios';

export default function App() {
  const [form, setForm] = useState({
    name: '',
    id: '',
    flightNumber: '',
    date: '',
    aircraft: 'ATR'
  });
  const [seats, setSeats] = useState([]);
  const [error, setError] = useState('');

  const handleChange = e => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const formatDate = (dateString) => {
    const date = new Date(dateString);
    const year = date.getFullYear();
    const month = String(date.getMonth() + 1).padStart(2, '0');
    const day = String(date.getDate()).padStart(2, '0');
    return `${year}-${month}-${day}`;
  };

  const handleSubmit = async () => {
    setError('');
    setSeats([]);
    
    // Validation
    if (!form.name || !form.id || !form.flightNumber || !form.date || !form.aircraft) {
      setError('All fields are required.');
      return;
    }

    const formattedDate = formatDate(form.date);
    try {
      const check = await axios.post('http://localhost:8080/api/check', {
        flightNumber: form.flightNumber,
        date: formattedDate
      });
      if (check.data.exists) {
        setError('Vouchers already generated for this flight and date.');
        return;
      }
      const gen = await axios.post('http://localhost:8080/api/generate', {
        ...form,
        date: formattedDate
      });
      if (gen.data.success) {
        setSeats(gen.data.seats);
      }
    } catch (err) {
      setError(err.response?.data?.error || 'Something went wrong');
    }
  };

  return (
      <div className="max-w-xl mx-auto p-6 space-y-4">
        <h1 className="text-2xl font-bold">Voucher Generator</h1>
        <div className="space-y-2">
          <div className="flex items-center space-x-2">
            <label className="w-32">Crew Name:</label>
            <input className="border p-2 flex-1" placeholder="Crew Name" name="name" onChange={handleChange} value={form.name} required />
          </div>
          <div className="flex items-center space-x-2">
            <label className="w-32">Crew ID:</label>
            <input className="border p-2 flex-1" placeholder="Crew ID" name="id" onChange={handleChange} value={form.id} required />
          </div>
          <div className="flex items-center space-x-2">
            <label className="w-32">Flight Number:</label>
            <input className="border p-2 flex-1" placeholder="Flight Number" name="flightNumber" onChange={handleChange} value={form.flightNumber} required />
          </div>
          <div className="flex items-center space-x-2">
            <label className="w-32">Flight Date:</label>
            <input type="date" className="border p-2 flex-1" name="date" onChange={handleChange} value={form.date} required />
          </div>
          <div className="flex items-center space-x-2">
            <label className="w-32">Aircraft Type:</label>
            <select className="border p-2 flex-1" name="aircraft" onChange={handleChange} value={form.aircraft} required>
              <option value="ATR">ATR</option>
              <option value="Airbus 320">Airbus 320</option>
              <option value="Boeing 737 Max">Boeing 737 Max</option>
            </select>
          </div>
          <div className="flex justify-end">
            <button onClick={handleSubmit} className="bg-blue-600 text-white px-4 py-2 rounded">
              Generate Vouchers
            </button>
          </div>
        </div>
        {error && <p className="text-red-500 mt-4">{error}</p>}
        {seats.length > 0 && (
            <div>
              <h2 className="font-bold mt-4">Assigned Seats:</h2>
              <ul className="list-disc pl-6">
                {seats.map(seat => <li key={seat}>{seat}</li>)}
              </ul>
            </div>
        )}
      </div>
  );
}

