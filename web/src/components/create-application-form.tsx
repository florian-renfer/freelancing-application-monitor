"use client";

import { CircleCheck, CircleX, Dock, Globe, PenLine, Send } from "lucide-react";

import {
  InputGroup,
  InputGroupAddon,
  InputGroupInput,
  InputGroupTextarea,
} from "@/components/ui/input-group";

import {
  Select,
  SelectContent,
  SelectGroup,
  SelectItem,
  SelectLabel,
  SelectTrigger,
  SelectValue,
} from "@/components/ui/select";
import { useState } from "react";
import { Datepicker } from "./datepicker";

export function CreateApplicationForm() {
  let [state, setState] = useState<string>("");

  return (
    <div className="grid w-full max-w-sm gap-6">
      <InputGroup>
        <InputGroupInput placeholder="Title..." />
        <InputGroupAddon>
          <Dock />
        </InputGroupAddon>
      </InputGroup>

      <InputGroup>
        <InputGroupInput placeholder="https://www.freelancermap.de/projekt/typo3-entwickler-fuer-updates-wartung-und-relaunch-gesucht" />
        <InputGroupAddon>
          <Globe />
        </InputGroupAddon>
      </InputGroup>

      <ApplicationStateSelect onValueChange={(val) => setState(val)} />

      {/* FIXME: this is currently not working because of nested button tags */}
      {/* {state && state !== "DRAFT" && <Datepicker />} */}

      <InputGroup>
        <InputGroupTextarea placeholder="Describe the project..." />
        <InputGroupAddon>
          <PenLine />
        </InputGroupAddon>
      </InputGroup>
    </div>
  );
}

export function ApplicationStateSelect({
  ...props
}: React.ComponentProps<typeof Select>) {
  return (
    <Select {...props}>
      <SelectTrigger className="w-full">
        <SelectValue placeholder="Select an Application state" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Application states</SelectLabel>
          <SelectItem value="DRAFT" className="flex gap-2">
            <PenLine />
            Draft
          </SelectItem>
          <SelectItem value="APPLIED" className="flex gap-2">
            <Send />
            Applied
          </SelectItem>
          <SelectItem value="HIRED" className="flex gap-2">
            <CircleCheck className="text-green-400" />
            Hired
          </SelectItem>
          <SelectItem value="REJECTED" className="flex gap-2 ">
            <CircleX className="text-red-400" />
            Rejected
          </SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  );
}
