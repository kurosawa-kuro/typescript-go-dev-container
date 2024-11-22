'use client';
import { useAuthStore } from '@/stores/authStore';

export default function Home() {
  const { user, isLoading, login, logout } = useAuthStore();

  const loginUser = {
    email: 'user1@example.com',
    password: 'password',
  };

  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <main className="container mx-auto px-4 py-8">
        <h1 className="text-2xl font-bold mb-4">State Management Debug</h1>
        
        <div className="bg-white dark:bg-gray-800 p-4 rounded-lg shadow">
          <h2 className="text-xl font-semibold mb-2">Auth Store State:</h2>
          <pre className="bg-gray-100 dark:bg-gray-700 p-2 rounded">
            {JSON.stringify({ user, isLoading }, null, 2)}
          </pre>

          <div className="mt-4 space-y-2">
            <button
              onClick={() => login(loginUser.email, loginUser.password)}
              className="bg-blue-500 text-white px-4 py-2 rounded mr-2"
              disabled={isLoading}
            >
              Test Login
            </button>
            
            <button
              onClick={logout}
              className="bg-red-500 text-white px-4 py-2 rounded"
              disabled={isLoading}
            >
              Test Logout
            </button>
          </div>
        </div>
      </main>
    </div>
  );
}
