-- Drop constraints
ALTER TABLE public_vote DROP CONSTRAINT unique_public_vote_election_puclic;
ALTER TABLE candidate DROP CONSTRAINT unique_candidate_election_puclic;

-- Drop tables in reverse order of dependencies
DROP TABLE IF EXISTS vote;
DROP TABLE IF EXISTS public_vote;
DROP TABLE IF EXISTS candidate;
DROP TABLE IF EXISTS election;