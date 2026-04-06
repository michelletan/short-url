"use client";
import { useState } from 'react';
import { useRouter } from 'next/navigation';

export default function UrlForm() {
  const [url, setUrl] = useState('');
  const [error, setError] = useState('');
  const router = useRouter();

  const handleSubmit = async (e) => {
    e.preventDefault();
    setError('');
    if (!/^https?:\/\/.+\..+/.test(url)) {
      setError('Please enter a valid URL.');
      return;
    }
    // Placeholder for API call
    const code = 'abc123'; // Replace with real API response
    router.push(`/success?code=${code}`);
  };

  return (
    <form onSubmit={handleSubmit}>
      <input
        type="url"
        placeholder="Enter your URL"
        value={url}
        onChange={e => setUrl(e.target.value)}
        required
      />
      <button type="submit">Shorten</button>
      {error && <p style={{ color: 'red' }}>{error}</p>}
    </form>
  );
}