'use client';

import Link from 'next/link';
import { usePathname } from 'next/navigation';

const navigation = [
  { name: 'Dev', href: '/dev' },
  { name: 'Health Check', href: '/dev/health-check' },
  { name: 'Login', href: '/dev/login' },
  { name: 'State Management', href: '/dev/state-management' },
  { name: 'PGAdmin', href: 'http://localhost:5050/' },
  { name: 'Storybook', href: 'http://localhost:6006/' },
  { name: 'Swagger', href: 'http://localhost:8000/swagger/index.html' },
  { name: 'Home', href: '/' },
  { name: 'Admin', href: '/admin' },
];

export function DevSidebar() {
  const pathname = usePathname();

  return (
    <div className="w-64 bg-gray-600 text-white min-h-screen p-4">
      <div className="mb-8">
        <h1 className="text-xl font-bold">Go Node.js App</h1>
      </div>
      
      <nav className="space-y-2">
        {navigation.map((item) => {
          const isActive = pathname === item.href;
          return (
            <Link
              key={item.name}
              href={item.href}
              className={`
                block px-4 py-2 rounded-lg transition-colors
                ${isActive 
                  ? 'bg-gray-800 text-white' 
                  : 'text-gray-300 hover:bg-gray-800 hover:text-white'
                }
              `}
            >
              {item.name}
            </Link>
          );
        })}
      </nav>
    </div>
  );
}
