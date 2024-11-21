import Image from "next/image";
import { ServerResponse } from "@/app/components/dev/ServerResponse"; 

export default function Home() {
  return (
    <div className="min-h-screen bg-gray-50 dark:bg-gray-900">
      <main className="container mx-auto px-4 py-8">
        <ServerResponse />
      </main>
    </div>
  );
}
