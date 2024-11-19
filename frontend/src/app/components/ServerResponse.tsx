async function getPingData() {
  try {
    const res = await fetch('http://backend:8000/ping', {
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });
  
    if (!res.ok) {
      throw new Error(`Failed to fetch data: ${res.status}`);
    }
  
    return res.json();
  } catch (error) {
    console.error('Error fetching data:', error);
    return { error: 'Failed to fetch data from server' };
  }
}

export async function ServerResponse() {
  const data = await getPingData();

  return (
    <div className="p-4 bg-gray-100 dark:bg-gray-800 rounded-lg">
      <h2 className="text-lg font-bold mb-2">Server Response:</h2>
      <pre className="bg-white dark:bg-gray-900 p-3 rounded">
        {JSON.stringify(data, null, 2)}
      </pre>
    </div>
  );
}