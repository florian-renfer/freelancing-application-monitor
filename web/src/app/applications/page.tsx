import { CreateApplicationForm } from "@/components/create-application-form";
import { Badge } from "@/components/ui/badge";

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

import {
  Item,
  ItemActions,
  ItemContent,
  ItemDescription,
  ItemGroup,
  ItemSeparator,
  ItemTitle,
} from "@/components/ui/item";

import { Separator } from "@/components/ui/separator";

import { Application } from "@/types/application";

import {
  BadgeCheckIcon,
  ChevronRightIcon,
  GalleryVerticalEnd,
  PlusCircleIcon,
} from "lucide-react";
import Link from "next/link";

import React from "react";

function EmptyState() {
  return (
    <Empty>
      <EmptyHeader>
        <EmptyMedia variant="icon">
          <GalleryVerticalEnd />
        </EmptyMedia>
        <EmptyTitle>No Applications Yet</EmptyTitle>
        <EmptyDescription>
          You haven&apos;t applied to any projects yet. Get started by adding
          your first application.
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
                  Keep track of a new Application. We are going to take care of
                  mandatory steps like <em>checking for duplicates</em>.
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
  );
}

interface ApplicationListProps {
  applications: Application[];
}

function ApplicationList({ applications }: ApplicationListProps) {
  return (
    <div className="flex w-full max-w-full flex-col gap-6">
      <ItemGroup>
        {applications.map((application, index) => (
          <React.Fragment key={application.id}>
            <Item asChild>
              <Link href={`/applications/${application.id}`}>
                <ItemContent className="gap-1">
                  <ItemTitle>
                    <Badge
                      variant={
                        application.state.toLowerCase() as
                          | "default"
                          | "secondary"
                          | "destructive"
                          | "outline"
                          | "applied"
                      }
                    >
                      <BadgeCheckIcon />
                      {application.state}
                    </Badge>
                    {application.title}
                  </ItemTitle>
                  <ItemDescription>{application.description}</ItemDescription>
                </ItemContent>
                <ItemActions>
                  <ChevronRightIcon className="size-4" />
                </ItemActions>
              </Link>
            </Item>
            {index !== applications.length - 1 && <ItemSeparator />}
          </React.Fragment>
        ))}
      </ItemGroup>
    </div>
  );
}

export default async function Applications() {
  const data = await fetch(process.env.API_BASE_URL + "/applications");
  const applications = (await data.json()) as Application[];

  return (
    <>
      <header>
        <div className="flex justify-between">
          <h1>Applications</h1>
          <Button>
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

      {applications && applications.length > 0 ? (
        <ApplicationList applications={applications} />
      ) : (
        <EmptyState />
      )}
    </>
  );
}
