import { DevSidebar } from "@/app/components/dev/DevSidebar";

export default function DevTemplate({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <div className="flex min-h-screen">
      <DevSidebar />
      <main className="flex-1 overflow-auto">
        {children}
      </main>
    </div>
  );
}