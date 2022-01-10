-- CreateTable
CREATE TABLE "User" (
    "id" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
    "username" TEXT NOT NULL,
    "email" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
);

-- CreateTable
CREATE TABLE "Profile" (
    "id" TEXT NOT NULL PRIMARY KEY,
    "userId" INTEGER NOT NULL,
    "name" TEXT NOT NULL,
    CONSTRAINT "Profile_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "Hour" (
    "uid" TEXT NOT NULL PRIMARY KEY,
    "userId" INTEGER NOT NULL,
    "task" TEXT NOT NULL,
    "locationUid" TEXT NOT NULL,
    "hours" REAL NOT NULL,
    "comment" TEXT NOT NULL,
    "created_at" DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT "Hour_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "Hour_locationUid_fkey" FOREIGN KEY ("locationUid") REFERENCES "Location" ("uid") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "Location" (
    "uid" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL
);

-- CreateTable
CREATE TABLE "Trail" (
    "uid" TEXT NOT NULL PRIMARY KEY,
    "name" TEXT NOT NULL,
    "locationUid" TEXT,
    CONSTRAINT "Trail_locationUid_fkey" FOREIGN KEY ("locationUid") REFERENCES "Location" ("uid") ON DELETE SET NULL ON UPDATE CASCADE
);

-- CreateTable
CREATE TABLE "HoursOnTrails" (
    "trailUid" TEXT NOT NULL,
    "hourUid" TEXT NOT NULL,

    PRIMARY KEY ("trailUid", "hourUid"),
    CONSTRAINT "HoursOnTrails_trailUid_fkey" FOREIGN KEY ("trailUid") REFERENCES "Trail" ("uid") ON DELETE RESTRICT ON UPDATE CASCADE,
    CONSTRAINT "HoursOnTrails_hourUid_fkey" FOREIGN KEY ("hourUid") REFERENCES "Hour" ("uid") ON DELETE RESTRICT ON UPDATE CASCADE
);

-- CreateIndex
CREATE UNIQUE INDEX "User_username_key" ON "User"("username");

-- CreateIndex
CREATE UNIQUE INDEX "User_email_key" ON "User"("email");

-- CreateIndex
CREATE UNIQUE INDEX "Profile_userId_key" ON "Profile"("userId");

-- CreateIndex
CREATE UNIQUE INDEX "Location_name_key" ON "Location"("name");
