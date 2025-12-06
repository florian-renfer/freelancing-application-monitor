import { Button } from "@/components/ui/button";
import clsx from "clsx";
import {
  ArrowLeftIcon,
  CalendarDaysIcon,
  CheckCircleIcon,
  CreditCardIcon,
  ExternalLinkIcon,
  FilePenIcon,
  PaperclipIcon,
  UserCircleIcon,
} from "lucide-react";
import Link from "next/link";
import { SlashIcon } from "lucide-react";

import {
  Breadcrumb,
  BreadcrumbItem,
  BreadcrumbLink,
  BreadcrumbList,
  BreadcrumbPage,
  BreadcrumbSeparator,
} from "@/components/ui/breadcrumb";
import { Application } from "@/types/application";
import { Badge } from "@/components/ui/badge";

const invoice = {
  subTotal: "$8,800.00",
  tax: "$1,760.00",
  total: "$10,560.00",
  items: [
    {
      id: 1,
      title: "Logo redesign",
      description: "New logo and digital asset playbook.",
      hours: "20.0",
      rate: "$100.00",
      price: "$2,000.00",
    },
    {
      id: 2,
      title: "Website redesign",
      description: "Design and program new company website.",
      hours: "52.0",
      rate: "$100.00",
      price: "$5,200.00",
    },
    {
      id: 3,
      title: "Business cards",
      description: 'Design and production of 3.5" x 2.0" business cards.',
      hours: "12.0",
      rate: "$100.00",
      price: "$1,200.00",
    },
    {
      id: 4,
      title: "T-shirt design",
      description: "Three t-shirt design concepts.",
      hours: "4.0",
      rate: "$100.00",
      price: "$400.00",
    },
  ],
};
const activity = [
  {
    id: 1,
    type: "created",
    person: { name: "Chelsea Hagon" },
    date: "7d ago",
    dateTime: "2023-01-23T10:32",
  },
  {
    id: 2,
    type: "edited",
    person: { name: "Chelsea Hagon" },
    date: "6d ago",
    dateTime: "2023-01-23T11:03",
  },
  {
    id: 3,
    type: "sent",
    person: { name: "Chelsea Hagon" },
    date: "6d ago",
    dateTime: "2023-01-23T11:24",
  },
  {
    id: 4,
    type: "commented",
    person: {
      name: "Chelsea Hagon",
      imageUrl:
        "https://images.unsplash.com/photo-1550525811-e5869dd03032?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80",
    },
    comment:
      "Called client, they reassured me the invoice would be paid by the 25th.",
    date: "3d ago",
    dateTime: "2023-01-23T15:56",
  },
  {
    id: 5,
    type: "viewed",
    person: { name: "Alex Curren" },
    date: "2d ago",
    dateTime: "2023-01-24T09:12",
  },
  {
    id: 6,
    type: "paid",
    person: { name: "Alex Curren" },
    date: "1d ago",
    dateTime: "2023-01-24T09:20",
  },
];

export function BreadcrumbWithCustomSeparator() {
  return (
    <Breadcrumb>
      <BreadcrumbList>
        <BreadcrumbItem>
          <BreadcrumbLink asChild>
            <Link href="/dashboard">Home</Link>
          </BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbSeparator>
          <SlashIcon />
        </BreadcrumbSeparator>
        <BreadcrumbItem>
          <BreadcrumbLink asChild>
            <Link href="/applications">Applications</Link>
          </BreadcrumbLink>
        </BreadcrumbItem>
        <BreadcrumbSeparator>
          <SlashIcon />
        </BreadcrumbSeparator>
        <BreadcrumbItem>
          <BreadcrumbPage>Details</BreadcrumbPage>
        </BreadcrumbItem>
      </BreadcrumbList>
    </Breadcrumb>
  );
}

export default async function ApplicationDetailsPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;

  const data = await fetch(`${process.env.API_BASE_URL}/applications/${id}`);
  const application = (await data.json()) as Application;

  return (
    <>
      <header className="relative isolate">
        <div className="py-10">
          <div className="pb-6 flex gap-4 items-center">
            <Button variant="secondary" asChild>
              <Link href="/applications">
                <ArrowLeftIcon className="size-4" />
                <span className="ml-2">Back</span>
              </Link>
            </Button>

            <BreadcrumbWithCustomSeparator />
          </div>
          <div className="mx-auto flex max-w-2xl items-center justify-between gap-x-8 lg:mx-0 lg:max-w-none">
            <div className="flex items-center gap-x-6">
              <img
                alt=""
                src="https://tailwindcss.com/plus-assets/img/logos/48x48/tuple.svg"
                className="size-16 flex-none rounded-full bg-gray-100 outline -outline-offset-1 outline-black/5 dark:bg-gray-800 dark:outline-white/10"
              />
              <h1>
                <div className="text-sm/6 text-gray-500 dark:text-gray-400">
                  Application{" "}
                  <span className="text-gray-700 dark:text-gray-300">
                    #{id}
                  </span>
                </div>
                <div className="mt-1 text-base font-semibold text-gray-900 dark:text-white">
                  {application.title}
                </div>
              </h1>
            </div>
            <div className="flex items-center gap-x-4">
              <Button variant="secondary" asChild>
                <Link
                  href={application.url}
                  target="_blank"
                  rel="noopener noreferrer"
                >
                  Open URL
                  <ExternalLinkIcon className="ml-2 size-4" />
                </Link>
              </Button>
              <Button asChild>
                <Link href={`/applications/${id}/edit`}>
                  Edit
                  <FilePenIcon className="ml-2 size-4" />
                </Link>
              </Button>
            </div>
          </div>
        </div>
      </header>

      <div className="mx-auto grid max-w-2xl grid-cols-1 grid-rows-1 items-start gap-x-8 gap-y-8 lg:mx-0 lg:max-w-none lg:grid-cols-3">
        {/* Invoice summary */}
        <div className="lg:col-start-3 lg:row-end-1">
          <h2 className="sr-only">Summary</h2>
          <div className="rounded-lg bg-gray-50 shadow-xs outline-1 outline-gray-900/5 dark:bg-gray-800/50 dark:shadow-none dark:-outline-offset-1 dark:outline-white/10">
            <dl className="flex flex-wrap">
              <div className="flex-auto pt-6 pl-6">
                <dt className="text-sm/6 font-semibold text-gray-900 dark:text-white">
                  Hourly Rate
                </dt>
                <dd className="mt-1 text-base font-semibold text-gray-900 dark:text-white">
                  € 130.00
                </dd>
              </div>
              <div className="flex-none self-end px-6 pt-4">
                <dt className="sr-only">Status</dt>
                <dd>
                  <Badge
                    variant={
                      application.state.toLowerCase() as "draft" | "applied"
                    }
                  >
                    {application.state}
                  </Badge>
                </dd>
              </div>
              <div className="mt-6 flex w-full flex-none gap-x-4 border-t border-gray-900/5 px-6 pt-6 dark:border-white/10">
                <dt className="flex-none">
                  <span className="sr-only">Client</span>
                  <UserCircleIcon
                    aria-hidden="true"
                    className="h-6 w-5 text-gray-400 dark:text-gray-500"
                  />
                </dt>
                <dd className="text-sm/6 font-medium text-gray-900 dark:text-white">
                  Renfer, Florian
                </dd>
              </div>
              <div className="mt-4 flex w-full flex-none gap-x-4 px-6">
                <dt className="flex-none">
                  <span className="sr-only">Applied on</span>
                  <CalendarDaysIcon
                    aria-hidden="true"
                    className="h-6 w-5 text-gray-400 dark:text-gray-500"
                  />
                </dt>

                {application.applied_at && (
                  <dd className="text-sm/6 text-gray-500 dark:text-gray-400">
                    <time dateTime={application.applied_at.toString()}>
                      {application.applied_at.toString()}
                    </time>
                  </dd>
                )}
              </div>
              <div className="mt-4 flex w-full flex-none gap-x-4 px-6">
                <dt className="flex-none">
                  <span className="sr-only">Status</span>
                  <CreditCardIcon
                    aria-hidden="true"
                    className="h-6 w-5 text-gray-400 dark:text-gray-500"
                  />
                </dt>
                <dd className="text-sm/6 text-gray-500 dark:text-gray-400">
                  Applied via{" "}
                  <Link
                    href={application.url}
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    {application.url}
                  </Link>
                </dd>
              </div>
            </dl>
            <div className="mt-6 border-t border-gray-900/5 px-6 py-6 dark:border-white/10">
              <a
                href="#"
                className="text-sm/6 font-semibold text-gray-900 dark:text-white"
              >
                Download receipt <span aria-hidden="true">&rarr;</span>
              </a>
            </div>
          </div>
        </div>

        {/* Invoice */}
        <div className="-mx-4 px-4 py-8 shadow-xs ring-1 ring-gray-900/5 sm:mx-0 sm:rounded-lg sm:px-8 sm:pb-14 lg:col-span-2 lg:row-span-2 lg:row-end-2 xl:px-16 xl:pt-16 xl:pb-20 dark:shadow-none dark:ring-white/10">
          <h2 className="text-base font-semibold text-gray-900 dark:text-white">
            Application
          </h2>
          <dl className="mt-6 grid grid-cols-1 text-sm/6 sm:grid-cols-2">
            <div className="sm:pr-4">
              <dt className="inline text-gray-500 dark:text-gray-400">
                Applied on
              </dt>{" "}
              {application.applied_at && (
                <dd className="inline text-gray-700 dark:text-gray-300">
                  <time dateTime={application.applied_at.toString()}>
                    {application.applied_at.toString()}
                  </time>
                </dd>
              )}
            </div>
            <div className="mt-2 sm:mt-0 sm:pl-4">
              <dt className="inline text-gray-500 dark:text-gray-400">
                Created on
              </dt>{" "}
              {application.created_at && (
                <dd className="inline text-gray-700 dark:text-gray-300">
                  <time dateTime={application.created_at.toString()}>
                    {application.created_at.toLocaleString()}
                  </time>
                </dd>
              )}
            </div>
            <div className="mt-6 border-t border-gray-900/5 pt-6 sm:pr-4 dark:border-white/10">
              <dt className="font-semibold text-gray-900 dark:text-white">
                From
              </dt>
              <dd className="mt-2 text-gray-500 dark:text-gray-400">
                <span className="font-medium text-gray-900 dark:text-white">
                  Acme, Inc.
                </span>
                <br />
                7363 Cynthia Pass
                <br />
                Toronto, ON N3Y 4H8
              </dd>
            </div>
            <div className="mt-8 sm:mt-6 sm:border-t sm:border-gray-900/5 sm:pt-6 sm:pl-4 dark:sm:border-white/10">
              <dt className="font-semibold text-gray-900 dark:text-white">
                To
              </dt>
              <dd className="mt-2 text-gray-500 dark:text-gray-400">
                <span className="font-medium text-gray-900 dark:text-white">
                  {application.title}{" "}
                </span>
                <br />
                886 Walter Street
                <br />
                New York, NY 12345
              </dd>
            </div>
          </dl>
          <table className="mt-16 w-full text-left text-sm/6 whitespace-nowrap">
            <colgroup>
              <col className="w-full" />
              <col />
              <col />
              <col />
            </colgroup>
            <thead className="border-b border-gray-200 text-gray-900 dark:border-white/15 dark:text-white">
              <tr>
                <th scope="col" className="px-0 py-3 font-semibold">
                  Projects
                </th>
                <th
                  scope="col"
                  className="hidden py-3 pr-0 pl-8 text-right font-semibold sm:table-cell"
                >
                  Hours
                </th>
                <th
                  scope="col"
                  className="hidden py-3 pr-0 pl-8 text-right font-semibold sm:table-cell"
                >
                  Rate
                </th>
                <th
                  scope="col"
                  className="py-3 pr-0 pl-8 text-right font-semibold"
                >
                  Price
                </th>
              </tr>
            </thead>
            <tbody>
              {invoice.items.map((item) => (
                <tr
                  key={item.id}
                  className="border-b border-gray-100 dark:border-white/10"
                >
                  <td className="max-w-0 px-0 py-5 align-top">
                    <div className="truncate font-medium text-gray-900 dark:text-white">
                      {item.title}
                    </div>
                    <div className="truncate text-gray-500 dark:text-gray-400">
                      {item.description}
                    </div>
                  </td>
                  <td className="hidden py-5 pr-0 pl-8 text-right align-top text-gray-700 tabular-nums sm:table-cell dark:text-gray-300">
                    {item.hours}
                  </td>
                  <td className="hidden py-5 pr-0 pl-8 text-right align-top text-gray-700 tabular-nums sm:table-cell dark:text-gray-300">
                    {item.rate}
                  </td>
                  <td className="py-5 pr-0 pl-8 text-right align-top text-gray-700 tabular-nums dark:text-gray-300">
                    {item.price}
                  </td>
                </tr>
              ))}
            </tbody>
            <tfoot>
              <tr>
                <th
                  scope="row"
                  className="px-0 pt-6 pb-0 font-normal text-gray-700 sm:hidden dark:text-gray-300"
                >
                  Subtotal
                </th>
                <th
                  scope="row"
                  colSpan={3}
                  className="hidden px-0 pt-6 pb-0 text-right font-normal text-gray-700 sm:table-cell dark:text-gray-300"
                >
                  Subtotal
                </th>
                <td className="pt-6 pr-0 pb-0 pl-8 text-right text-gray-900 tabular-nums dark:text-white">
                  {invoice.subTotal}
                </td>
              </tr>
              <tr>
                <th
                  scope="row"
                  className="pt-4 font-normal text-gray-700 sm:hidden dark:text-gray-300"
                >
                  Tax
                </th>
                <th
                  scope="row"
                  colSpan={3}
                  className="hidden pt-4 text-right font-normal text-gray-700 sm:table-cell dark:text-gray-300"
                >
                  Tax
                </th>
                <td className="pt-4 pr-0 pb-0 pl-8 text-right text-gray-900 tabular-nums dark:text-white">
                  {invoice.tax}
                </td>
              </tr>
              <tr>
                <th
                  scope="row"
                  className="pt-4 font-semibold text-gray-900 sm:hidden dark:text-white"
                >
                  Total
                </th>
                <th
                  scope="row"
                  colSpan={3}
                  className="hidden pt-4 text-right font-semibold text-gray-900 sm:table-cell dark:text-white"
                >
                  Total
                </th>
                <td className="pt-4 pr-0 pb-0 pl-8 text-right font-semibold text-gray-900 tabular-nums dark:text-white">
                  {invoice.total}
                </td>
              </tr>
            </tfoot>
          </table>
        </div>

        <div className="lg:col-start-3">
          {/* Activity feed */}
          <h2 className="text-sm/6 font-semibold text-gray-900 dark:text-white">
            Activity
          </h2>
          <ul role="list" className="mt-6 space-y-6">
            {activity.map((activityItem, activityItemIdx) => (
              <li key={activityItem.id} className="relative flex gap-x-4">
                <div
                  className={clsx(
                    activityItemIdx === activity.length - 1
                      ? "h-6"
                      : "-bottom-6",
                    "absolute top-0 left-0 flex w-6 justify-center",
                  )}
                >
                  <div className="w-px bg-gray-200 dark:bg-white/10" />
                </div>
                {activityItem.type === "commented" ? (
                  <>
                    <img
                      alt=""
                      src={activityItem.person.imageUrl}
                      className="relative mt-3 size-6 flex-none rounded-full bg-gray-100 outline -outline-offset-1 outline-black/5 dark:bg-gray-800 dark:outline-white/10"
                    />
                    <div className="flex-auto rounded-md p-3 ring-1 ring-gray-200 ring-inset dark:ring-white/10">
                      <div className="flex justify-between gap-x-4">
                        <div className="py-0.5 text-xs/5 text-gray-500 dark:text-gray-400">
                          <span className="font-medium text-gray-900 dark:text-white">
                            {activityItem.person.name}
                          </span>{" "}
                          commented
                        </div>
                        <time
                          dateTime={activityItem.dateTime}
                          className="flex-none py-0.5 text-xs/5 text-gray-500 dark:text-gray-400"
                        >
                          {activityItem.date}
                        </time>
                      </div>
                      <p className="text-sm/6 text-gray-500 dark:text-gray-400">
                        {activityItem.comment}
                      </p>
                    </div>
                  </>
                ) : (
                  <>
                    <div className="relative flex size-6 flex-none items-center justify-center bg-white dark:bg-gray-900">
                      {activityItem.type === "paid" ? (
                        <CheckCircleIcon
                          aria-hidden="true"
                          className="size-6 text-indigo-600 dark:text-indigo-500"
                        />
                      ) : (
                        <div className="size-1.5 rounded-full bg-gray-100 ring-1 ring-gray-300 dark:bg-white/10 dark:ring-white/20" />
                      )}
                    </div>
                    <p className="flex-auto py-0.5 text-xs/5 text-gray-500 dark:text-gray-400">
                      <span className="font-medium text-gray-900 dark:text-white">
                        {activityItem.person.name}
                      </span>{" "}
                      {activityItem.type} the invoice.
                    </p>
                    <time
                      dateTime={activityItem.dateTime}
                      className="flex-none py-0.5 text-xs/5 text-gray-500 dark:text-gray-400"
                    >
                      {activityItem.date}
                    </time>
                  </>
                )}
              </li>
            ))}
          </ul>

          {/* New comment form */}
          <div className="mt-6 flex gap-x-3">
            <img
              alt=""
              src="https://images.unsplash.com/photo-1472099645785-5658abf4ff4e?ixlib=rb-1.2.1&ixid=eyJhcHBfaWQiOjEyMDd9&auto=format&fit=facearea&facepad=2&w=256&h=256&q=80"
              className="size-6 flex-none rounded-full bg-gray-100 outline -outline-offset-1 outline-black/5 dark:bg-gray-800 dark:outline-white/10"
            />
            <form action="#" className="relative flex-auto">
              <div className="overflow-hidden rounded-lg pb-12 outline-1 -outline-offset-1 outline-gray-300 focus-within:outline-2 focus-within:-outline-offset-2 focus-within:outline-indigo-600 dark:bg-white/5 dark:outline-white/10 dark:focus-within:outline-indigo-500">
                <label htmlFor="comment" className="sr-only">
                  Add your comment
                </label>
                <textarea
                  id="comment"
                  name="comment"
                  rows={2}
                  placeholder="Add your comment..."
                  className="block w-full resize-none bg-transparent px-3 py-1.5 text-base text-gray-900 placeholder:text-gray-400 focus:outline-none sm:text-sm/6 dark:text-white dark:placeholder:text-gray-500"
                  defaultValue={""}
                />
              </div>

              <div className="absolute inset-x-0 bottom-0 flex justify-between py-2 pr-2 pl-3">
                <div className="flex items-center space-x-5">
                  <Button variant="ghost" className="rounded-full">
                    <PaperclipIcon aria-hidden="true" className="size-4" />
                    <span className="sr-only">Attach a file</span>
                  </Button>
                </div>

                <Button type="submit" variant="secondary">
                  Comment
                </Button>
              </div>
            </form>
          </div>
        </div>
      </div>
    </>
  );
}
