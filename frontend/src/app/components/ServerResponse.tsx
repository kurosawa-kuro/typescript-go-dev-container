async function getPingData() {
  try {
    const res = await fetch('http://app:8000/ping', {
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });
  
    if (!res.ok) {
      throw new Error(`Failed to fetch ping data: ${res.status}`);
    }
  
    return res.json();
  } catch (error) {
    console.error('Error fetching ping data:', error);
    return { error: 'Failed to fetch ping data from server' };
  }
}

async function getMicroposts() {
  try {
    const res = await fetch('http://app:8000/microposts', {
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch microposts: ${res.status}`);
    }

    return res.json();
  } catch (error) {
    console.error('Error fetching microposts:', error);
    return { error: 'Failed to fetch microposts from server' };
  }
}

export async function ServerResponse() {
  const pingData = await getPingData();
  const microposts = await getMicroposts();

  return (
    <div className="p-4 bg-gray-100 dark:bg-gray-800 rounded-lg">
      <h2 className="text-lg font-bold mb-2">Server Response:</h2>
      <div className="space-y-4">
        <div>
          <h3 className="font-semibold mb-1">Ping Response:</h3>
          <pre className="bg-white dark:bg-gray-900 p-3 rounded">
            {JSON.stringify(pingData, null, 2)}
          </pre>
        </div>
        <div>
          <h3 className="font-semibold mb-1">Microposts:</h3>
          <pre className="bg-white dark:bg-gray-900 p-3 rounded">
            {JSON.stringify(microposts, null, 2)}
          </pre>
        </div>
      </div>
    </div>
  );
}