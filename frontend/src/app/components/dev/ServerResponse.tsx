type ApiResponse<T> = {
  data?: T;
  error?: string;
};

type PingResponse = {
  message: string;
};

type PingDBResponse = {
  message: string;
};

type Micropost = {
  id: number;
  title: string;
  created_at: string;
  updated_at: string;
};

async function fetchFromApi<T>(endpoint: string): Promise<ApiResponse<T>> {
  try {
    const res = await fetch(`http://app:8000${endpoint}`, {
      headers: {
        'Content-Type': 'application/json',
      },
      cache: 'no-store',
    });
  
    if (!res.ok) {
      throw new Error(`Failed to fetch data: ${res.status}`);
    }
  
    const data = await res.json();
    return { data };
  } catch (error) {
    console.error(`Error fetching from ${endpoint}:`, error);
    return { error: `Failed to fetch data from ${endpoint}` };
  }
}

function ResponseSection({ title, data }: { title: string; data: any }) {
  const isPingResponse = title.includes('Ping');
  
  return (
    <div className="w-full">
      <h3 className="text-xl font-semibold mb-3 text-gray-800 dark:text-gray-200">
        {title}
      </h3>
      <pre className={`bg-white dark:bg-gray-800 p-6 rounded-lg shadow-sm overflow-auto text-sm
        ${isPingResponse ? 'min-h-[150px]' : 'min-h-[600px]'}`}>
        {JSON.stringify(data, null, 2)}
      </pre>
    </div>
  );
}

export async function ServerResponse() {
  const pingResponse = await fetchFromApi<PingResponse>('/ping');
  const pingDBResponse = await fetchFromApi<PingDBResponse>('/ping-db');
  const micropostsResponse = await fetchFromApi<Micropost[]>('/microposts');

  return (
    <div className="bg-gray-100 dark:bg-gray-800/50 rounded-xl p-8 shadow-lg">
      <h2 className="text-2xl font-bold mb-6 text-gray-900 dark:text-gray-100">
        Server Response
      </h2>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-8">
        <ResponseSection 
          title="Ping Response" 
          data={pingResponse.data || pingResponse.error} 
        />
        <ResponseSection 
          title="Ping DB Response" 
          data={pingDBResponse.data || pingDBResponse.error} 
        />
        <ResponseSection 
          title="Microposts" 
          data={micropostsResponse.data || micropostsResponse.error} 
        />
      </div>
    </div>
  );
}