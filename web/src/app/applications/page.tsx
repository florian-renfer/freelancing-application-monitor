import { CreateApplicationForm } from "@/components/create-application-form";
import { Button } from "@/components/ui/button";
import {
  Dialog,
  DialogClose,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger,
} from "@/components/ui/dialog";
import {
  Empty,
  EmptyHeader,
  EmptyMedia,
  EmptyTitle,
  EmptyDescription,
  EmptyContent,
} from "@/components/ui/empty";
import { Input } from "@/components/ui/input";
import { Separator } from "@/components/ui/separator";
import { Label } from "@radix-ui/react-label";
import { GalleryVerticalEnd, PlusCircleIcon } from "lucide-react";

export default function Applications() {
  return (
    <>
      <header>
        <div className="flex justify-between">
          <h1>Applications</h1>
          <Button variant="secondary">
            <PlusCircleIcon />
            Create Application
          </Button>
        </div>
        <p className="text-muted-foreground">
          Manage all of your applications in one place. Create new ones or
          update existing ones.
        </p>
        <Separator className="mt-4" />
      </header>

      <div className="flex items-center justify-center h-full">
        <Empty>
          <EmptyHeader>
            <EmptyMedia variant="icon">
              <GalleryVerticalEnd />
            </EmptyMedia>
            <EmptyTitle>No Applications Yet</EmptyTitle>
            <EmptyDescription>
              You haven&apos;t applied to any projects yet. Get started by
              adding your first application.
            </EmptyDescription>
          </EmptyHeader>
          <EmptyContent>
            <Dialog>
              <form>
                <DialogTrigger asChild>
                  <Button variant="secondary" className="hover:cursor-pointer">
                    <PlusCircleIcon />
                    Create Application
                  </Button>
                </DialogTrigger>
                <DialogContent className="sm:max-w-[425px]">
                  <DialogHeader>
                    <DialogTitle>Create Application</DialogTitle>
                    <DialogDescription>
                      Keep track of a new Application. We are going to take care
                      of mandatory steps like <em>checking for duplicates</em>.
                    </DialogDescription>
                  </DialogHeader>
                  <CreateApplicationForm />
                  <DialogFooter>
                    <DialogClose asChild>
                      <Button variant="outline">Cancel</Button>
                    </DialogClose>
                    <Button type="submit">Create Application</Button>
                  </DialogFooter>
                </DialogContent>
              </form>
            </Dialog>
          </EmptyContent>
        </Empty>
      </div>
    </>
  );
}
